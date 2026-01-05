# Go Todo CLI

A simple and clean command-line todo application written in Go.
Supports adding, editing, deleting, toggling, and listing todos with persistent storage.

---

## ✨ Features

- Add, edit, delete, and toggle todos
- Persistent storage using JSON file
- Generic storage layer using Go generics
- Clean CLI interface using `flag`
- Unit tests and CI with GitHub Actions

---

## 📦 Requirements

- Go 1.21+

Check your Go version:

```bash
go version
```

---

## 🚀 Usage

Clone the repository:

```bash
git clone https://github.com/Kaveh-Goodarzi/go-todo-cli.git
cd go-todo-cli
```

Run the application:

```bash
go run main.go [flags]
```

---

## 🔧 Available Commands

| Flag | Description |
|-----|-------------|
| `-add "title"` | Add a new todo |
| `-edit id:title` | Edit a todo |
| `-del id` | Delete a todo |
| `-toggle id` | Toggle todo completion |
| `-list` | List all todos |


**Example**:

```bash
go run main.go -add "Learn Go"
go run main.go -list
go run main.go -toggle 0
```

---

## 🧪 Running Tests

```bash
go test -v
```

---

## 🗂 Project Structure

```
    CLI-ToDo-App/
    ├── .github/workflows/go.yml
    ├── main.go
    ├── todo.go
    ├── storage.go
    ├── command.go
    ├── main_test.go
    ├── go.mod
    └── README.md
```

---


## 🤖 Continuous Integration

This project uses GitHub Actions to automatically run tests on every push and pull request.