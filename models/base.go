package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := "postgres"
	password := "postgres"
	dbName := "the_ledger"
	dbHost := "localhost"
	dbPort := "5432"

	conn, err := gorm.Open("postgres", "host="+dbHost+" user="+username+" dbname="+dbName+" password="+password+" port="+dbPort)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
