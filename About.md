# Task Tracker CLI

Task Tracker CLI is a lightweight command-line project developed as an implementation of the [Task Tracker](https://roadmap.sh/projects/task-tracker) backend challenge from roadmap.sh. 

## The Challenge Context
The roadmap.sh project requires building a simple CLI application to track tasks, focusing on core programming fundamentals:
- **File System Operations**: Reading and writing directly to a local JSON file (`tasks.json`).
- **Command-line Parsing**: Handling user inputs via positional arguments.
- **Data Structures**: Managing tasks with properties like `id`, `description`, `status` (`todo`, `in-progress`, `done`), `createdAt`, and `updatedAt`.

## Our Implementation
This implementation is written natively in Go. Although the challenge guidelines suggest minimizing external frameworks, we utilized the robust Go ecosystem (such as `cobra` for elegant CLI routing) while adhering precisely to the functional requirements of the challenge.

### Fulfilled Requirements
- [x] Add, Update, and Delete tasks
- [x] Mark a task as `in-progress` or `done`
- [x] List all tasks
- [x] List tasks filtered by status (e.g., list `done`, list `todo`)
- [x] Native JSON file storage (auto-created if it doesn't exist)
- [x] Error and edge-case handling

## Goal
This project serves as a practical demonstration of building clean, maintainable, and portable CLI tools using Go, directly translating a well-structured set of requirements into a functional product.
