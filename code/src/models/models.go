package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	dbhost = "192.168.0.36"
	dbport = "32768"
	dbuser = "postgres"
	dbpass = "mysecretpassword"
	dbname = "postgres"
)

func dbConfig() map[string]string {
	conf := make(map[string]string)

	conf[dbhost] = dbhost
	conf[dbport] = dbport
	conf[dbuser] = dbuser
	conf[dbpass] = dbpass
	conf[dbname] = dbname
	return conf
}

var db *sql.DB

func InitDB() *sql.DB {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport], config[dbuser], config[dbpass], config[dbname])
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Success")
	return db
}
