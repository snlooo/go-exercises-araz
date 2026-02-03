# Homework: Session 4 - TaskManagementSystem

## Objective
In this session, you will refactor the `TaskManagementSystem` solution from Session 3 by implementing Go's function declarations, multiple return values, recursion, variadic functions, and handling errors using `defer`, `panic`, and `recover`.

## Submission
- Create a new branch named `SES-04/functions`, and submit all deliverables to this branch in your private GitHub Repository.
- Create a Pull Request (PR) and send the link to the instructor for review.
- At last your repository structure should be like below:

  ```bash
  ├───session-1
  ├───session-2
  ├───session-3
  ├───session-4
  └───TaskManagementSystem
          go.mod
          main.go
  ```

- The deadline for submission is **next session date - 1 day**.



## Instructions:

### **Instructions:**

1.  **Refactor the code into functions**:
    -   Move the logic for task completion checks and task remaining calculations into separate functions.
    -   Move task listing also to separate function
    -   Define and use functions to print task statuses.

2.  **Implement multiple return values**:
    -   Create a function that returns both the `tasksRemaining` and a boolean indicating if the project is near completion (e.g., if more than 90 tasks are completed).

3.  **Use recursion**:
    -   Implement a recursive function to simulate a countdown that shows how many tasks are left to be completed, reducing the count one by one.

4.  **Variadic functions**:
    -   Implement a variadic function to accept and print the details of any number of tasks (e.g., task names).

5.  **Defer, panic, and recover**:
    -   Create a scenario where you use `defer` to clean up after checking the project status.
    -   Simulate a `panic` when the `tasksCompleted` exceeds the total tasks, and use `recover` to handle this case gracefully.


5. Run the program to verify that `based on cases` everything is working correctly.

    ```
      Recovered from panic: Error: tasksCompleted exceeds total tasks
      Status check completed.
      Task list:
      1: Task 1
      2: Task 2
      3: Task 3
      4: Task 4
      5: Task 5
      Recursive countdown of remaining tasks:
      Tasks remaining: 95
      Tasks remaining: 94
      Tasks remaining: 93
      ...
      Tasks remaining: 1
      All tasks completed!
      Is the project completed? false
    ```

**Test:**
- Change `completed task` to 105 and check the recovered message from panic
- Change `completed task` to 100 and check `project status`
- Change `completed task` to 95 and check `project status`, `phase`

**Deliverable:** 
- Submit your task.