#  Session 2 Homework

## Objective
To reinforce the concepts of Go's data types, variables, constants, and formatting learned in the second session.

## Submission
- Create a new branch named `SES-02/basic-syntax-data-types`, and submit all deliverables to this branch in your private GitHub Repository.
- Create a Pull Request (PR) and send the link to the instructor for review.
- At last your repository structure should be like below:

  ```bash
  ├───session-2
          go.mod
          main.go
  ```

- The deadline for submission is **next session date - 1 day**.

## Instructions:
1. Extend the existing project: In the same project created in session-1, add new functionality in the `main.go` file to introduce and work with variables, constants, and various data types.

2. Add the following code:

    - Variables:
        - Declare and initialize variables representing:
            - The current project status (e.g., "active", "in progress").
            - The number of tasks created so far (integer).
            - A flag to indicate whether the project is completed (boolean).
    - Constants:
        - Declare a constant for the total number of available tasks (e.g., 100).
        - Use iota to declare constants representing different task priorities (e.g., low, medium, high).

3. Display the variables and constants:

    - Print the values of the above variables and constants in a formatted way using the fmt package.
    - Format the output to display:
        - Project status.
        - Tasks created so far out of total tasks.
        - Task priorities (use iota to assign priority levels).
        - Project completion status.

4. Type conversion and string formatting:

    - Convert the boolean value (project completed) into a string and display it in the output message.


5. Run the program to verify that everything is working correctly.

    ```
      Welcome to the Task Management System!
      Project start date is: 2024-09-18 00:00:00
      Project: Task Management System

      Current project status: IN PROGRESS
      Tasks completed: 25 out of 100
      Task priorities: 0-Low, 1-Medium, 2-High
      Is the project completed? false
    ```


**Deliverable:**
- Submit your task.




# References
- **Mandatory** [Strings in Golang](https://www.programiz.com/golang/string)
- **Mandatory:** [fmt.Printf() Function in Golang With Examples](https://www.geeksforgeeks.org/fmt-printf-function-in-golang-with-examples/)
- **Mandatory:** [strconv package in Golang](https://www.geeksforgeeks.org/strconv-package-in-golang/)
- **Optional:** [ASCII Table](https://www.ascii-code.com/)