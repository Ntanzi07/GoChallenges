# Challenge: Task Management API (TODO API)

## Objective

Create a **REST API in Go using Fiber** that allows you to **manage tasks** — list, create, update, and delete.

---

## Description

The API must support the following operations:

| Method | Route | Description |
|--------|--------|-------------|
| `GET` | `/tasks` | Returns all tasks |
| `GET` | `/tasks/:id` | Returns a specific task |
| `POST` | `/tasks` | Creates a new task |
| `PUT` | `/tasks/:id` | Updates an existing task |
| `DELETE` | `/tasks/:id` | Deletes a task |

---

## Project Structure

```
/todo-api
├── main.go
├── go.mod
├── data/
│   └── tasks.go
├── handlers/
│   └── tasks.go
├── models/
│   └── task.go
└── routes/
    └── routes.go
```

---

## Data Model

```go
type Task struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"created_at"`
}
```

---

## Requirements

- Use **Fiber** (`go get github.com/gofiber/fiber/v2`).
- Store data **in memory** (a slice `[]Task`).
- Each new task should automatically receive a new `ID` and `CreatedAt = time.Now()`.

---

## Extra Challenge (optional)

- Persist data in a `tasks.json` file.
- Add validation for POST (do not accept empty titles).
- Filter tasks by status: `/tasks?completed=true` or `/tasks?completed=false`.

---

## Example Requests

### POST `/tasks`
**Body:**
```json
{
  "title": "Study Go and Fiber"
}
```
**Response:**
```json
{
  "id": 1,
  "title": "Study Go and Fiber",
  "completed": false,
  "created_at": "2025-10-31T14:23:00Z"
}
```

---

### GET `/tasks`
```json
{
  "tasks": [
    {
      "id": 1,
      "title": "Study Go and Fiber",
      "completed": false,
      "created_at": "2025-10-31T14:23:00Z"
    }
  ]
}
```

---

### PUT `/tasks/1`
**Body:**
```json
{
  "completed": true
}
```

---

### DELETE `/tasks/1`
**Response:**
```json
{
  "message": "Task successfully removed"
}
```

---

## Concepts Practiced

- Routing with Fiber  
- Structs and JSON  
- Slice manipulation and ID search  
- Layered architecture (`models`, `data`, `handlers`, `routes`)  
- Go project best practices  

---

### Tip

To run the API:

```bash
go mod init todo-api
go get github.com/gofiber/fiber/v2
go run main.go
```

---

Developed as a practical challenge to practice **Go + Fiber**.
