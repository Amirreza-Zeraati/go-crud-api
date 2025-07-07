# Go CRUD Service

A RESTful API for basic CRUD operations on a User model, built with Go and Gin.

## 📦 Features

- Create, Read, Update, Delete users
- Clean code structure with controllers and initializers

## 🛠 Tech Stack

- Go (Golang)
- Gin framework
- Gorm
- PostgreSQL | sqlite
- Environment configuration via `.env`

---

## 🚀 Endpoints

| Method | Endpoint     | Description          |
|--------|--------------|----------------------|
| POST   | /user        | Create a user        |
| GET    | /user        | List all users       |
| GET    | /user/:id    | Get a single user    |
| PUT    | /user/:id    | Update a user        |
| DELETE | /user/:id    | Delete a user        |

---

## ⚙️ Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL
- `.env` file in the root directory

### Example `.env`
```bash
PORT=3000
DB_URL=postgres://user:password@localhost:5432/dbname
```

---

### Run the Server

```bash
go run main.go
```


## 🧪 Testing
You can use tools like [Postman](https://www.postman.com/) to interact with the CRUD endpoints.
