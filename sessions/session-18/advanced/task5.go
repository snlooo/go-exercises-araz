package advanced

import (
	"fmt"
	"session-18/db_sql_pkg"
)

//Task 5: Using Prepared Statements
//
//Write a Go program that:
//Inserts a new student into the students table using a prepared statement.
//Prints "Insert successful" after the operation.

func AddNewStudent() {

	db, err := db_sql_pkg.ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	stmt, err := db.Prepare("INSERT INTO students VALUES ($1, $2 , $3)")

	_, err = stmt.Exec(4, "Araz", 25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Insert successful")
	}
}
