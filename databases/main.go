package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Actor struct {
	Id        int32
	Firstname string
	Lastname  string
}

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

	fmt.Println(strings.Repeat("-", 20))

	fmt.Printf("Retrive actor with ID: %d\n", actorId)

	actors, err := getActor(int32(actorId))
	if err != nil {
		log.Fatal(err)
	}

	for _, actor := range actors {
		fmt.Printf("Actor found: %v\n", actor)
	}
}
