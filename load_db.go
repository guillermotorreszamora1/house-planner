package main

import (
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func readDefault() {
	db, _ := sql.Open("sqlite3", "test_design")
	readTypes(db)
	readSubtypes(db)
	readUnits(db)
	readSupply(db)
}
func readTypes(db *sql.DB) {
	//Open the file default_db_values and print its' data
	file, _ := os.Open("default_db_values/type.csv")
	defer file.Close()

	reader := csv.NewReader(file)
	//Drop first row
	record, _ := reader.Read()
	if record[0] != "type_id" || record[1] != "type" {
		log.Fatal("type headers don't match")
		return
	}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		db.Exec("INSERT INTO type(type_id, type) VALUES (?, ?)", record[0], record[1])

	}
}
func readSubtypes(db *sql.DB) {
	file, _ := os.Open("default_db_values/subtype.csv")
	defer file.Close()

	reader := csv.NewReader(file)
	record, _ := reader.Read()
	if len(record) != 2 {
		log.Fatalf("subtype header length doesn't match record:%s", record)
	}
	if record[0] != "subtype_id" || record[1] != "subtype" {
		log.Fatalf("subtype headers don't match:%s %s", record[0], record[1])
		return
	}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		db.Exec("INSERT INTO subtype(subtype_id, subtype) VALUES (?, ?)", record[0], record[1])
	}
}

func readUnits(db *sql.DB) {
	file, _ := os.Open("default_db_values/units.csv")
	defer file.Close()

	reader := csv.NewReader(file)
	record, _ := reader.Read()
	if record[0] != "unit_type_id" || record[1] != "unit_name" {
		log.Fatal("unit headers don't match")
		return
	}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		db.Exec("INSERT INTO units(unit_type_id, unit_name) VALUES (?, ?)", record[0], record[1])
	}
}

type Supply struct {
	name         string
	type_id      int
	subtype_id   *int
	unit_type_id int
}

func validateType(db *sql.DB, typeName string) (bool, *int) {
	query, err := db.Query("SELECT type_id FROM type WHERE type = ?", typeName)
	if err != nil {
		return false, nil
	}
	var typeId int
	query.Next()
	err = query.Scan(&typeId)
	if err != nil {
		log.Fatalf("type doesn't exist |%s| err |%s|", typeName, err)
		return false, nil
	}
	return true, &typeId
}
func validateSubtype(db *sql.DB, subtypeName string) (bool, *int) {
	if subtypeName == "NULL" {
		return true, nil
	}
	query, err := db.Query("SELECT subtype_id FROM subtype WHERE subtype = ?", subtypeName)
	if err != nil {
		return false, nil
	}
	var typeId int
	query.Next()
	err = query.Scan(&typeId)
	if err != nil {
		log.Fatalf("subtype doesn't exist %s", subtypeName)
		return false, nil
	}
	return true, &typeId
}
func validateUnit(db *sql.DB, unitName string) (bool, *int) {
	query, err := db.Query("SELECT unit_type_id FROM units WHERE unit_name = ?", unitName)
	if err != nil {
		return false, nil
	}
	var unitId int
	query.Next()
	err = query.Scan(&unitId)
	if err != nil {
		log.Fatal("unit doesn't exist: ")
		return false, nil
	}
	return true, &unitId
}
func validateSuply(db *sql.DB, record []string) (bool, *Supply) {
	name := record[0]
	log.Printf("validating supply for |%s|", name)
	status, typeId := validateType(db, record[1])
	if status != true {
		log.Fatalf("type doesn't exist for %s", name)
		return false, nil
	}
	status, subTypeId := validateSubtype(db, record[2])
	if status != true {
		log.Fatalf("subtype doesn't exist for %s", name)
		return false, nil

	}
	status, unitType := validateUnit(db, record[3])
	if status != true {
		log.Fatalf("unit doesn't exist for %s", name)
		return false, nil
	}
	var supply Supply = Supply{name: name, type_id: *typeId, subtype_id: subTypeId, unit_type_id: *unitType}
	return true, &supply
}
func readSupply(db *sql.DB) {
	file, _ := os.Open("default_db_values/Supply.csv")
	defer file.Close()

	reader := csv.NewReader(file)
	//Drop first row
	record, _ := reader.Read()
	if record[0] != "name" || record[1] != "type" || record[2] != "subtype" || record[3] != "unit_type" {
		log.Fatal("supply headers don't match")
		return
	}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		status, supply := validateSuply(db, record)
		if status != true {
			log.Fatal("supply isn't valid")
			return
		}
		db.Exec("INSERT INTO Supply(name,type_id,subtype_id,unit_type_id) "+
			"VALUES (?,?,?,?)",
			supply.name,
			supply.type_id,
			supply.subtype_id,
			supply.unit_type_id)
	}
}
