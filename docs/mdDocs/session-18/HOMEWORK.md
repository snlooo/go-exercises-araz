# Session 18 Homework: Working with Databases

## Submission

- Create a new branch named `SES-18/db`, and submit all deliverables to this branch in your private GitHub Repository.
- Create a Pull Request (PR) and send the link to the instructor for review.
- At last your repository structure should be like below:

```go
  ├───session-19
         task
         ───sql_intro
              task1.go
              task2.go
         ───db_sql_pkg
              task3.go
              task4.go
         ───advanced
              task5.go
              task6.go
         go.mod
         main.go
```

- Run `go mod init session-18` in your `session-19` root folder
- In the `main.go` call all task's individually
- The deadline for submission is `next session date - 1 day`.

---

### **Topic 1: Introduction to SQL in Go**

**Task 1: Setting Up a PostgreSQL Environment with Docker**

1.  Write the Docker command to create and run a PostgreSQL container.
2.  Use the following parameters:
    - Container Name: `school-db`
    - Username: `student`
    - Password: `securepass`
    - Database Name: `school`
    - Port Mapping: `5432:5432`

**Expected Output:** Executing the command:

`docker run --name school-db -e POSTGRES_USER=student -e POSTGRES_PASSWORD=securepass -p 5432:5432 -d postgres`

Results in:

`<Container ID>`

Accessing the database using `DataGrip`, `Goland` or another tool

---

**Task 2: Basic SQL Commands**

1.  After setting up the PostgreSQL database, write SQL queries to:
    - Create a `students` table with columns `id (serial primary key)`, `name (text)`, and `age (int)`.
    - Insert three rows of data into the table.
    - Select all rows from the `students` table.

**Expected Output:** SQL commands:

```sql
CREATE TABLE students (id SERIAL PRIMARY KEY, name TEXT, age INT);
INSERT INTO students (name, age) VALUES ('Ali', 20), ('Geray', 22), ('Medina', 21);
SELECT * FROM students;
```

```
id |  name     | age
---+-----------+-----
1  | Ali       |  20
2  | Geray     |  22
3  | Medina    |  21
```

---

### **Topic 2: Using the `database/sql` Package**

**Task 3: Establishing a Database Connection**

1.  Write a Go program to connect to the `school` database using the `database/sql` package and the PostgreSQL driver.
2.  The program should:
    - Open a database connection.
    - Print "Connection successful" if the connection is valid.

**Expected Output:** :

`Connection successful`

---

**Task 4: Querying the Database**

1.  Extend the program from **Task 3** to:
    - Query all rows from the `students` table.
    - Print each student's `id`, `name`, and `age`.

**Expected Output:** :

```
ID: 1, Name: Ali, Age: 20
ID: 2, Name: Geray, Age: 22
ID: 3, Name: Medina, Age: 21
```

---

### **Topic 3: Advanced Database Interactions**

**Task 5: Using Prepared Statements**

1.  Write a Go program that:
    - Inserts a new student into the `students` table using a prepared statement.
    - Prints "Insert successful" after the operation.

**Expected Output:** :

`Insert successful`

---

**Task 6: Transactions in Go**

1.  Write a Go program that:
    - Starts a transaction.
    - Updates the age of `Ali` to `21`.
    - If successful, commits the transaction. If not, rolls back.

**Expected Output:**:

`Transaction successful`

---

## References

- [database/sql package](https://pkg.go.dev/database/sql)
- [PostgreSql tutorial](https://neon.tech/postgresql/tutorial)
- [SQL Drivers](https://go.dev/wiki/SQLDrivers)
- [database/sql tutorial](http://go-database-sql.org/overview.html)
