package db

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

func initDB() *gorm.DB {

	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		"127.0.0.1",
		"5432",
		"postgres",
		"eshop",
		"12345")
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatal("Couldn't connect to database", err.Error())
	}

	return db
}

// StartDbConnection Creates connection to database
func StartDbConnection() {
	database = initDB()
}

// GetDBConn func for getting db conn globally
func GetDBConn() *gorm.DB {
	return database
}
