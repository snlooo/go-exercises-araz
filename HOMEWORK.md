# Homework Assignment: Go Environment Setup and Introduction to Golang

## Objective
The goal of this homework is to get you familiar with the basics of Go, setting up your development environment, and writing your first simple Go program. Tasks should be placed under `session-1` folder. Coding homework always will be in the project named `TaskManagementSystem` and every change should push to the new branch for review.

## Submission
- Read [`GITHUB-USE.md`](https://github.com/zauradigozalov/mastering-golang-sessions/blob/main/GITHUB-USE.md) file carefully and create personal GitHub Repository.
- Create a new branch named `SES-01/first-project`, and submit all deliverables to this branch in your private GitHub Repository.
- Create a Pull Request (PR) and send the link to the instructor for review.
- At last your repository structure should be like below:

  ```bash
  ├───session-1
  │   ├───task-1
  │   │       screenshot.png
  │   │
  │   └───task-2
  │           screenshot.png
  │
  └───TaskManagementSystem
          go.mod
          main.go
  ```

- The deadline for submission is **next session date - 1 day**.

## Task 1: Setting Up Go Environment

### 1. Install Go
- Follow the installation guide provided on the official Go website [here](https://go.dev/doc/install).
- Confirm your installation by running `go version` in your terminal.

**Deliverable:** 
- Take a screenshot of your terminal showing the successful installation of Go and the output of the `go version` command. Place screenshot under `session-1\task-1` folder

### 2. Install GoLand
- Download and install GoLand IDE from [JetBrains](https://www.jetbrains.com/go/).
- When prompted, choose the **Free 30-day Trial** option if you don't have a license. This will allow you to use all features of GoLand during the trial period.

**Deliverable:** 
- Take a screenshot of your GoLand IDE and place under `session-1\task-1` folder

## Task 2: Basic Go Program and Code Formatting

### 1. Write a Simple Go Program
- Create a Go project under `TaskManagementSystem` folder:
  - Click on **File** > **New** > **Project**.
  - Choose **Go** as the project type. Choose latest Go version.
- Create a Go program `main.go` that prints below messages to the console
    ```
      Welcome to the Task Management System!
      Project start date is: 2024-09-18 00:00:00
      Project: Task Management System
    ```
- Use the `fmt` package to show messages
- Use the `gofmt` tool to format your Go program according to the Go style guidelines.
- Run `gofmt -w main.go` to apply the formatting changes.
- Run the program by clicking the green **Run** button.

**Deliverable:** 
- Submit your task.



# References
- **Mandatory:** [Git and GitHub Desktop for Beginners](https://www.youtube.com/watch?v=8Dd7KRpKeaE)
- **Mandatory:** [Review README.MD of course repository](https://github.com/zauradigozalov/mastering-golang-sessions)
- **Mandatory:** [Go Tour](https://go.dev/tour/welcome/1)
- **Optional:** [Go Playground](https://go.dev/play/)
- **Optional:** [Go Proverbs](https://go-proverbs.github.io/)
- **Optional:** [Awesome Gophers](https://awesome-go.com/gophers/)