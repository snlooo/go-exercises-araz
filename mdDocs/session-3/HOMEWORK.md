# Homework Assignment: TaskManagementSystem - Homework for Session 3

## Objective
Extend the TaskManagementSystem with more advanced Go features such as operators, conditional statements, loops, and error handling.

## Submission
- Create a new branch named `SES-03/operators-and-control-structures`, and submit all deliverables to this branch in your private GitHub Repository.
- Create a Pull Request (PR) and send the link to the instructor for review.
- At last your repository structure should be like below:

  ```bash
  ├───session-1
  ├───session-2
  ├───session-3
  └───TaskManagementSystem
          go.mod
          main.go
  ```

- The deadline for submission is **next session date - 1 day**.



## Instructions:

1. **Operators**:
   - Use arithmetic operators to dynamically calculate how many tasks remain.
   - Use logical operators in combination with conditions to check if the project is near completion (e.g., more than 90 tasks completed).

2. **Conditional Statements**:
   - Add an `if-else` block to check if the project is completed based on the number of tasks completed. If completed, update the status.
   - Add a `switch` statement that categorizes the project progress:
     - If less than 30 tasks are completed, mark the progress as "Starting phase".
     - If 30 to 60 tasks are completed, mark it as "Midway".
     - If more than 60 tasks are completed, mark it as "Final phase".

3. **Loops**:
   - Use a `for` loop to iterate tasks by number. For simplicity, create a for loop and iterate the remaining number of tasks, and output with task names (e.g., `"Task 1"`, `"Task 2"`, etc.).

4. **Error Handling**:
   - Implement a simple error-checking mechanism: if the tasks completed exceeds the total number of tasks, print an error message and reset the tasks completed to the maximum number.
   - Simulate an error scenario where the project cannot proceed if no tasks are completed. Use a custom error and handle it gracefully.


5. Run the program to verify that everything is working correctly.

    ```
      Welcome to the Task Management System!
      Project start date is: 2024-09-18 00:00:00 +0000 UTC 
      Project: Task Management System                      
                                                          
      Tasks remaining 5 out of 100                         
      Current project status: IN PROGRESS                  
      Project is in the final phase.                       
      Task priorities: 1-Low, 2-Medium, 3-High             
      Task list:                                           
      - Task 1                                             
      - Task 2                                             
      - Task 3                                             
      - Task 4                                             
      - Task 5                                             
      Is the project completed? false 
    ```

**Test:**
- Change `completed task` to 25 and check `project status`, `phase`, `task list`
- Change `completed task` to 95 and check `project status`, `phase`, `task list`

**Deliverable:** 
- Submit your task.




# References
- **Mandatory** [Go operators](https://www.codecademy.com/resources/docs/go/operators)
- **Mandatory:** [Golang: Labels](https://ravichaganti.com/blog/get-set-go-labels-in-go-language/)