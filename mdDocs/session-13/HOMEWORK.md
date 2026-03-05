# Homework: File Handling and I/O

## Submission
- Create a new branch named `SES-13/file-handling-and-io`, and submit all deliverables to this branch in your private GitHub Repository.
- Create a Pull Request (PR) and send the link to the instructor for review.
- At last your repository structure should be like below:

```go
  ├───session-1
  ├───session-2
  ├───session-3
  ├───session-4
  ├───session-5
  ├───session-6
  ├───session-7
  ├───session-8  
  ├───session-9
  ├───session-10
  ├───session-11
  ├───session-12
  ├───session-13
         task
         ───reading_writing
              task1.go
              task2.go
         ───json
              task3.go
              task4.go
         ───xml
              task5.go
              task6.go
         go.mod
         main.go
  └───TaskManagementSystem
```

- Run `go mod init session-13` in your `session-13` root folder
- In the `main.go` call all task's individually
- The deadline for submission is `next session date - 1 day`.

* * * * *
### Topic 1: File Reading

**Task 1: Read and Process CSV Data**

1.  Create a CSV file named `data.csv` with the following content:

```csv
name,age,grade
Alice,20,85
Bob,22,55
Carol,21,70
```

2.  Write a program that reads this file, processes each line, and prints only the names of students whose grade is above 60.

**Expected Output:**

```go
Students with passing grades:
- Alice
- Carol
```

**Task 2: Count Lines in a Text File**

1.  Create a text file named `story.txt` with several lines of text, such as:

```
Once upon a time, there was a brave knight.
He fought dragons and saved villages.
The people loved him dearly.
```

2.  Write a program that reads this file and counts the total number of lines.

3.  Print the line count.

**Expected Output:**

```
Total lines in file: 3
```

* * * * *

### Topic 2: Working with JSON

**Task 3: Read JSON Data and Filter by Condition**

1.  Create a JSON file named `students.json` with an array of student objects. Each student should have fields like `name`, `age`, and `grade`.
2.  Write a program that reads this file, filters out students with grades below a specified threshold (e.g., grade < 60), and prints the names of the remaining students.

**Sample `students.json`:**

```json
[
{"name": "Ali", "age": 20, "grade": 47},
{"name": "Sabina", "age": 22, "grade": 55},
{"name": "Medina", "age": 21, "grade": 85}
]
```

**Expected Output:**

```
Students with passing grades:
- Ali
- Sabina
```

**Task 4: JSON Configuration File Reader**

1.  Create a configuration JSON file, `config.json`, with fields such as `app_name`, `version`, `debug_mode`, and an array of `services` (e.g., `["service1", "service2"]`).
2.  Write a program that reads the JSON file and prints out the configuration details in a formatted manner.

**Sample `config.json`:**

```json
{
"app_name": "MyApp",
"version": "1.0",
"debug_mode": true,
"services": ["service1", "service2"]
}
```

**Expected Output:**

```
Configuration Details:
- App Name: MyApp
- Version: 1.0
- Debug Mode: Enabled
- Services: service1, service2
```

* * * * *

### Topic 3: Working with XML

**Task 5: XML Reader and Data Extraction**

1.  Create an XML file named `employees.xml` containing a list of employees. Each employee should have `name`, `position`, and `salary` fields.
2.  Write a program that reads the XML file and extracts only those employees whose salary is above a specified threshold (e.g., salary > 50000).

**Sample `employees.xml`:**

```xml
<employees>
<employee>
<name>Natig</name>
<position>Manager</position>
<salary>75000</salary>
</employee>
<employee>
<name>Subhan</name>
<position>Developer</position>
<salary>45000</salary>
</employee>
</employees>
```

**Expected Output:**

```
Employees with Salary above 50000:
- Natig, Manager
```

**Task 6: XML Config Writer**

1.  Create a program that generates an XML configuration file named `config.xml` with fields like `database`, `username`, `password`, and a nested `options` field that includes `auto_backup` and `max_connections`.
2.  Write the XML structure to `config.xml` and ensure it's formatted in a readable way.

**Expected Output in `config.xml`:**

```xml
<config>
<database>mydb</database>
<username>admin</username>
<password>secret</password>
<options>
<auto_backup>true</auto_backup>
<max_connections>100</max_connections>
</options>
</config>
```

## References
-[An overview of os/io packages](https://reintech.io/blog/an-overview-of-gos-os-and-io-packages)
-[IO package](https://pkg.go.dev/io)
-[OS package](https://pkg.go.dev/os)