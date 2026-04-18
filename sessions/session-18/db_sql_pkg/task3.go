package db_sql_pkg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Task 3: Establishing a Database Connection
//
//Write a Go program to connect to the school database using the database/sql package and the PostgreSQL driver.
//The program should:
//Open a database connection.
//Print "Connection successful" if the connection is valid.

func ConnectDB() (*sql.DB, error) {

	connString := "host=localhost port=5432 user=student password=securepass dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("DB connection failed:", err)
		return nil, err
	}

	fmt.Println("Connection successful")

	return db, nil
}
