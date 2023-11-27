package handlers

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func sqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := "go_test"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "go_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		fmt.Println("Database failed!")
		return nil
	}

	fmt.Println("Database connected!")
	return db
}
