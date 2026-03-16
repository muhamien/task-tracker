# Task Tracker CLI

A simple Command-Line Interface (CLI) application built with Go to track and manage your daily tasks. It stores tasks in a local `tasks.json` file. This project is a solution to the [Task Tracker](https://roadmap.sh/projects/task-tracker) backend challenge from roadmap.sh.

## Features

- Add a new task
- Update an existing task
- Delete a task
- Mark a task's status (`todo`, `in-progress`, `done`)
- List all tasks (or filter them by status)

## Setup & Installation

Make sure you have Go installed on your machine.

1.  **Clone the Repository**
    ```bash
    git clone git@github.com:muhamien/task-tracker.git
    cd task-tracker
    ```

2.  **Build the Project**
    ```bash
    go build -o task-tracker main.go
    ```

3.  **Run the CLI**
    ```bash
    ./task-tracker <command> [arguments]
    ```

## Usage

### 1. Add a Task
```bash
./task-tracker add "Learn Go"
```

### 2. Update a Task
```bash
./task-tracker update 1 "Learn Go and write tests"
```

### 3. Delete a Task
```bash
./task-tracker delete 1
```

### 4. Mark Task as In Progress
```bash
./task-tracker mark-in-progress 1
```

### 5. Mark Task as Done
```bash
./task-tracker mark-done 1
```

### 6. List Tasks
List all tasks:
```bash
./task-tracker list
```

## Repository Push Issue

When pushing via SSH (`git@github.com`), port 22 is sometimes blocked by local firewalls or the environment.
If you face connection issues, consider these alternatives:

1. **Use HTTPS**:
   ```bash
   git remote set-url origin https://github.com/muhamien/task-tracker.git
   git push -u origin main
   ```

2. **Route SSH over port 443**: Add the following to `~/.ssh/config`:
   ```ssh-config
   Host github.com
     Hostname ssh.github.com
     Port 443
     User git
   ```
