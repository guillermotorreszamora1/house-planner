package main

import (
	"database/sql"
	"log"
)

func showDb() {
	db, err := sql.Open("sqlite3", "test_design")
	if err != nil {
		log.Fatal(err)
	}
	showTypeTable(db)
	showSubTypeTable(db)
	showUnitsTable(db)
	db.Close()
}
func showTypeTable(db *sql.DB) {
	rows, err := db.Query("Select type_id, type from type")
	if err != nil {
		return
	}
	log.Printf("Printing type table")
	for rows.Next() {
		var typeId int
		var typeName string
		if err := rows.Scan(&typeId, &typeName); err != nil {
			log.Fatal(err)
		}
		log.Printf("id: %d name: %s\n", typeId, typeName)
	}
}
func showSubTypeTable(db *sql.DB) {
	rows, err := db.Query("Select subtype_id, subtype from subtype")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Printing subtype table")
	for rows.Next() {
		var subtypeId int
		var subtypeName string
		if err := rows.Scan(&subtypeId, &subtypeName); err != nil {
			log.Fatal(err)
		}
		log.Printf("id %d name is %s\n", subtypeId, subtypeName)
	}
}
func showUnitsTable(db *sql.DB) {
	rows, err := db.Query("Select unit_id, unit from unit")
	if err != nil {
		return
	}
	log.Printf("Printing unit table")
	for rows.Next() {
		var unitId int
		var unitName string
		if err := rows.Scan(&unitId, &unitName); err != nil {
			log.Fatal(err)
		}
		log.Printf("id: %d name: %s\n", unitId, unitName)

	}
}
