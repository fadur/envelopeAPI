package main

import (
	models "envelopeApi/code/src/models"
	"fmt"
)

func main() {
	fmt.Println("testing connection ..")
	db, err := models.InitDB()
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(models.Login{})
	db.AutoMigrate(models.Session{})
}
