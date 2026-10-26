package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const usage = `Usage:

  mycli <command> [arguments]

Commands:
  status          Show service status
  echo <number>   Echo a number
  help            Show this help

Options:
  --help, -h      Show help
`

const statusUsage = `Usage:
  mycli status [--help]

Show service status.

Options:
  --help, -h      Show this help
`

func main() {
	if len(os.Args) < 2 {
		fmt.Print(usage)
		return
	}

	//Syntax is a mixture of rust an C
	switch os.Args[1] {
	case "--help", "-h":
		fmt.Print(usage)
	case "status":
		fmt.Print("Status OK")
	case "echo":
		if len(os.Args) != 3 {
			fmt.Print(usage)
			return
		}
		fmt.Println(os.Args[2])

	case "show_db":
		showDb()
	case "load_data":
		readDefault()
		showDb()
	case "rest_api":
		db, _ := sql.Open("sqlite3", "test_design")
		defer db.Close()
		restApi(db)
	}
}
