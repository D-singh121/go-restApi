package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//----global variable for database connection

var Database *gorm.DB

var urlDSN = "user:Password@@tcp(localhost:3306)/myDB?parseTime=true"
var err error

func DataMigration() {
	//database connection basic code
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Connection failed :")
	}
	Database.AutoMigrate(&Employee{})

}
