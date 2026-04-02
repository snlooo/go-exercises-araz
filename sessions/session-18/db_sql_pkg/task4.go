package db_sql_pkg

import "fmt"

//Task 4: Querying the Database
//
//Extend the program from Task 3 to:
//Query all rows from the students table.
//Print each student's id, name, and age.

func PrintStudents() {
	db, err := ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("select * from students")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

}
