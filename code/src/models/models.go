package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	dbhost = "10.0.1.21"
	// dbhost = "192.168.0.36"
	dbport = "32768"
	dbuser = "postgres"
	dbpass = "mysecretpassword"
	dbname = "envelope"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpass, dbname)
	var err error
	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Success")

	db.LogMode(true)
	return db, err
}
