# 🧩 Desafio: API de Gerenciamento de Tarefas (TODO API)

## 🎯 Objetivo

Criar uma **API REST em Go usando Fiber** que permita **gerenciar tarefas** — listar, criar, atualizar e excluir.

---

## 📜 Descrição

A API deve permitir as seguintes operações:

| Método | Rota | Descrição |
|--------|-------|------------|
| `GET` | `/tasks` | Retorna todas as tarefas |
| `GET` | `/tasks/:id` | Retorna uma tarefa específica |
| `POST` | `/tasks` | Cria uma nova tarefa |
| `PUT` | `/tasks/:id` | Atualiza uma tarefa existente |
| `DELETE` | `/tasks/:id` | Remove uma tarefa |

---

## 🧱 Estrutura do Projeto

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

## 💾 Modelo de Dados

```go
type Task struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"created_at"`
}
```

---

## 🔧 Requisitos

- Use o **Fiber** (`go get github.com/gofiber/fiber/v2`).
- Os dados devem ser armazenados **em memória** (um slice `[]Task`).
- Cada nova tarefa recebe automaticamente um novo `ID` e `CreatedAt = time.Now()`.

---

## 🔥 Desafio Extra (opcional)

- Persistir dados em um arquivo `tasks.json`.
- Adicionar validações no POST (não aceitar título vazio).
- Filtrar tarefas por status: `/tasks?completed=true` ou `/tasks?completed=false`.

---

## 🧪 Exemplo de Requisições

### ➕ POST `/tasks`
**Body:**
```json
{
  "title": "Estudar Go e Fiber"
}
```
**Response:**
```json
{
  "id": 1,
  "title": "Estudar Go e Fiber",
  "completed": false,
  "created_at": "2025-10-31T14:23:00Z"
}
```

---

### 🔍 GET `/tasks`
```json
{
  "tasks": [
    {
      "id": 1,
      "title": "Estudar Go e Fiber",
      "completed": false,
      "created_at": "2025-10-31T14:23:00Z"
    }
  ]
}
```

---

### ✅ PUT `/tasks/1`
**Body:**
```json
{
  "completed": true
}
```

---

### ❌ DELETE `/tasks/1`
**Response:**
```json
{
  "message": "Tarefa removida com sucesso"
}
```

---

## 🎯 Conceitos Praticados

- Rotas no Fiber  
- Structs e JSON  
- Manipulação de slices e busca por ID  
- Organização em camadas (`models`, `data`, `handlers`, `routes`)  
- Boas práticas em projetos Go  

---

### 💡 Dica

Para rodar a API:

```bash
go mod init todo-api
go get github.com/gofiber/fiber/v2
go run main.go
```

---

👨‍💻 Desenvolvido como desafio prático para treinar **Go + Fiber**.
