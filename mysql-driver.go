package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func init() {
	var err error
	dsn := "root:Abdullah2007*@tcp(127.0.0.1:3306)/javatechie"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("Cannot connect to MySQL: %v", err)
	}
}
