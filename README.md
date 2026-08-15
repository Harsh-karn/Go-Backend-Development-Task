# Go - Backend Development Task

A production-ready REST API built with Go for managing users. The API stores user data (`name`, `dob`) in a PostgreSQL database and dynamically calculates their `age` when requested.

## 🚀 Features
- **Clean Architecture:** Divided into handlers, services, repositories, and models.
- **Dynamic Age Calculation:** Accurate age calculated on the fly without storing it in the DB.
- **Database Access:** Generated using `sqlc` for typesafe database queries.
- **Logging:** Structured logging implemented via Uber Zap.
- **Validation:** Input payload validation using `go-playground/validator`.
- **Middleware:** Request ID injection and request duration logging.
- **Pagination:** Supported on the list users endpoint.
- **Dockerized:** Ready to run with Docker Compose (App + Postgres).

## 🛠️ Tech Stack
- **Go 1.21**
- **GoFiber** (Web Framework)
- **PostgreSQL** (Database)
- **SQLC** (Code Generator for SQL)
- **Uber Zap** (Logger)
- **go-playground/validator** (Data Validation)

---

## 🏃 Getting Started

### Prerequisites
- Docker & Docker Compose
- (Optional) Go 1.21 if running locally outside Docker
- (Optional) sqlc if you wish to modify queries and regenerate code

### Running with Docker Compose (Recommended)
This will start both the PostgreSQL database and the API server.

```bash
docker-compose up --build
```
The API will be available at `http://localhost:3000`.

### Running Locally (Without Docker)

1. Start a PostgreSQL database instance and create a database named `userdb`.
2. Apply the migration in `db/migrations/000001_create_users_table.up.sql`.
3. Set your environment variables (or use defaults):
   ```bash
   export DB_DSN="postgres://user:password@localhost:5432/userdb?sslmode=disable"
   export PORT="3000"
   ```
4. Run the server:
   ```bash
   go run cmd/server/main.go
   ```

### Running Tests
To run the unit tests (e.g., the dynamic age calculation logic):
```bash
go test ./internal/service/... -v
```

---

## 🔌 API Endpoints

### 1. Create a User
**POST** `/users`

**Request:**
```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

**Response:** (`201 Created`)
```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}
```

### 2. Get User by ID
**GET** `/users/:id`

**Response:** (`200 OK`)
```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 36
}
```

### 3. Update a User
**PUT** `/users/:id`

**Request:**
```json
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

**Response:** (`200 OK`)
```json
{
  "id": 1,
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

### 4. Delete a User
**DELETE** `/users/:id`

**Response:** (`204 No Content`)

### 5. List All Users (with Pagination)
**GET** `/users?page=1&limit=10`

**Response:** (`200 OK`)
```json
[
  {
    "id": 1,
    "name": "Alice Updated",
    "dob": "1991-03-15",
    "age": 35
  }
]
```

---

## 📁 Project Structure

```text
/cmd/server/main.go       # Entry point of the application
/config/                  # Configuration loading (env vars)
/db/migrations/           # SQL migration files
/db/sqlc/                 # Auto-generated SQLC code
/db/query/                # SQL queries used by SQLC
/internal/
├── handler/              # HTTP handlers handling requests/responses
├── repository/           # Database abstraction layer
├── service/              # Core business logic (age calculation)
├── routes/               # API route definitions
├── middleware/           # Fiber middlewares (logging, request id)
├── models/               # Request/Response DTOs & struct validation
└── logger/               # Uber Zap logger setup
```
