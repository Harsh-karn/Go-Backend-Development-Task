# Go Backend Development Task

A RESTful API built with Go for managing users. The API stores user data (name, dob) in a PostgreSQL database and dynamically calculates their age when requested.

## Features

- **Architecture:** Layered design with handlers, services, repositories, and models.
- **Dynamic Calculation:** Age is calculated at runtime without being persisted in the database.
- **Database Access:** Typesafe database queries generated using SQLC.
- **Logging:** Structured logging implemented with Uber Zap.
- **Validation:** Input payload validation using go-playground/validator.
- **Middleware:** Custom middleware for request ID injection and request duration logging.
- **Pagination:** Supported on the list users endpoint.

## Tech Stack

- Go 1.21
- GoFiber
- PostgreSQL
- SQLC
- Uber Zap
- go-playground/validator

## Getting Started

### Running with Docker Compose (Recommended)

To start both the PostgreSQL database and the API server, run:

```bash
docker-compose up --build
```

The database migrations will run automatically on startup. The API will be available at `http://localhost:3000`.

### Running Locally (Without Docker)

1. Start a local PostgreSQL database instance and create a database named `userdb`.
2. Apply the migration located in `db/migrations/000001_create_users_table.up.sql`.
3. Set your environment variables (or use the defaults):
   ```bash
   export DB_DSN="postgres://user:password@localhost:5432/userdb?sslmode=disable"
   export PORT="3000"
   ```
4. Run the server:
   ```bash
   go run cmd/server/main.go
   ```

### Running Tests

To execute the unit tests for the dynamic age calculation logic:

```bash
go test ./internal/service/... -v
```

## API Endpoints

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

### 5. List All Users
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

## Project Structure

```text
/cmd/server/main.go       # Entry point
/config/                  # Configuration loading
/db/migrations/           # SQL migration files
/db/sqlc/                 # Auto-generated SQLC code
/db/query/                # SQL queries used by SQLC
/internal/
├── handler/              # HTTP handlers
├── repository/           # Database abstraction layer
├── service/              # Core business logic
├── routes/               # API route definitions
├── middleware/           # Fiber middlewares
├── models/               # Request/Response DTOs
└── logger/               # Uber Zap logger setup
```
