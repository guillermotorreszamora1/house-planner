package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func restApi(db *sql.DB) {
	r := gin.Default()

	// Allow all origins (local development)
	r.Use(cors.Default())

	// Define a simple GET endpoint
	r.GET("/ping", ping)
	r.GET("/types", listTypes(db))
	r.PUT("/types", addType(db))
	r.GET("/subtypes", listSubTypes(db))
	r.PUT("/subtypes", addSubType(db))
	r.GET("/units", listUnits(db))
	r.PUT("/units", addUnit(db))
	err := r.Run()
	if err != nil {
		return
	}
}
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

type SupplyType struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func listTypes(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("Select type_id, type from type")
		if err != nil {
			return
		}
		defer rows.Close()

		var supplyTypes []SupplyType
		for rows.Next() {

			var supplyType SupplyType
			if err := rows.Scan(&supplyType.ID, &supplyType.Name); err != nil {
				log.Fatal(err)
				return
			}
			supplyTypes = append(supplyTypes, supplyType)
		}
		c.JSON(http.StatusOK, supplyTypes)
	}
}
func addType(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var supplyType SupplyType
		err := c.ShouldBindJSON(&supplyType)
		if err != nil {
			logger.Info(err)
			c.AbortWithStatus(400)
			return
		}
		logger.Infof("Values obtained are |%s| and |%s|", supplyType.ID, supplyType.Name)
		_, err = db.Exec("INSERT INTO type(type_id, type) VALUES (?, ?)", supplyType.ID, supplyType.Name)
		if err != nil {
			logger.Info(err)
			c.AbortWithStatus(400)
			return
		}
		logger.Info("New type added to the DB")
	}
}

type SupplySubType struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func listSubTypes(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("Select subtype_id, subtype from subtype")
		if err != nil {
			return
		}
		defer rows.Close()

		var supplySubTypes []SupplySubType
		for rows.Next() {
			var supplySubType SupplySubType
			if err := rows.Scan(&supplySubType.ID, &supplySubType.Name); err != nil {
				log.Fatal(err)
				return
			}

			supplySubTypes = append(supplySubTypes, supplySubType)
		}
		c.JSON(http.StatusOK, supplySubTypes)
	}
}

func addSubType(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var supplySubType SupplySubType
		err := c.ShouldBindJSON(&supplySubType)
		if err != nil {
			logger.Info(err)
			c.AbortWithStatus(400)
			return
		}
		log.Printf("Values obtained are |%s| |%s|", supplySubType.ID, supplySubType.Name)
		_, err = db.Exec("INSERT INTO subtype(subtype_id, subtype) VALUES (?,?)", supplySubType.ID, supplySubType.Name)
		if err != nil {
			logger.Info(err)
			c.AbortWithStatus(400)
			return
		}
		logger.Info("New subtype added to the DB")
	}
}

type Units struct {
	ID   int    `json:"id" binding:"required"`   //SQL unit_type_id
	Name string `json:"name" binding:"required"` //unit_name
}

func listUnits(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("Select unit_type_id, unit_name from units")
		if err != nil {
			return
		}
		defer rows.Close()
		var unitTypes []Units
		for rows.Next() {
			var unitType Units
			if err := rows.Scan(&unitType.ID, &unitType.Name); err != nil {
				log.Fatal(err)
				return
			}
			unitTypes = append(unitTypes, unitType)
		}
		c.JSON(http.StatusOK, unitTypes)
	}
}
func addUnit(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var unitType Units
		err := c.ShouldBindJSON(&unitType)
		if err != nil {
			logger.Info(err)
			return
		}
		log.Printf("Values obtained are |%s| |%s|", unitType.ID, unitType.Name)
		_, err = db.Exec("INSERT INTO units(unit_type_id, unit_name) VALUES (?,?)", unitType.ID, unitType.Name)
		if err != nil {
			logger.Info(err)
			c.AbortWithStatus(400)
			return
		}
		logger.Info("New unit added to the DB")
	}
}

//Show supplies
