package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

mcr.microsoft.com/devcontainers/go:1-1.21.3-bullseye
var db *sql.DB

func main() {

	dsn := mysql.Config{
		User:   "josh",
		Passwd: "lupin3",
		Net:    "tcp",
		Addr:   "172.20.0.3:3306",
		DBName: "mydb",
	}

	var err error

	db, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)

	}

	fmt.Println("Connected!")

	actorId, err := addActor("Giorgio", "Ghost")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error inserting actor: %v", err))
	}

	fmt.Printf("Actor ID: %v\n", actorId)
}

func addActor(firstname, lastname string) (int64, error) {
	result, err := db.Exec("INSERT INTO actor (firstname, lastname) VALUES (?, ?)", firstname, lastname)
	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addActor Insert: %v", err)
	}

	return id, nil
}
