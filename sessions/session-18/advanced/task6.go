package advanced

import (
	"fmt"
	"log"
	"session-18/db_sql_pkg"
)

//Task 6: Transactions in Go
//
//Write a Go program that:
//Starts a transaction.
//Updates the age of Ali to 21.
//If successful, commits the transaction. If not, rolls back.

func UpdateStudentWithTransaction() {

	db, err := db_sql_pkg.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = tx.Exec("UPDATE  students SET name = $1 where id = $2", "Araz", 4)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	_, err = tx.Exec("UPDATE  students SET age = $1 where id = $2", 35, 6)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Transaction successful")

}
