package sql_intro

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//Task 1: Setting Up a PostgreSQL Environment with Docker
//
//Write the Docker command to create and run a PostgreSQL container.
//Use the following parameters:
//Container Name: school-db
//Username: student
//Password: securepass
//Database Name: school
//Port Mapping: 5432:5432

func ConnectDB() {
	connString := "host=localhost port=5432 user=student password=securepass dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("DB connection failed:", err)
	}

	fmt.Println("Connected successfully!")
}
