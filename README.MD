# Authentication API

RESTful Authentication API built with **Go (Golang)**, **Gin**, and **PostgreSQL**.  
This project implements a simple authentication system using a layered architecture (Handler → Service → Repository) with manual Dependency Injection.

## Features

- User Registration
- User Login
- Password hashing using bcrypt
- PostgreSQL database integration
- Database migration
- Layered Architecture
- Manual Dependency Injection

---

# Architecture

```
cmd/main.go                 → Application entry point

internal/
├── di/                     → Dependency Injection
├── handlers/               → HTTP Handlers
├── lib/                    → Response helpers
├── models/                 → Models & Request DTO
├── repo/                   → Data access layer
└── service/                → Business logic

migrations/
└── 000001_create_users_table
```

---

# Database Schema

```mermaid
erDiagram

USERS {
    bigint id PK
    string email UK
    string password
    string username
    string phone
    timestamp created_at
    timestamp updated_at
}
```

---

# Tech Stack

| Dependency | Purpose |
|------------|---------|
| Go | Programming Language |
| Gin | HTTP Web Framework |
| PostgreSQL | Database |
| pgx/v5 | PostgreSQL Driver |
| bcrypt | Password Hashing |
| godotenv | Environment Loader |

---

# Prerequisites

- Go 1.24+
- PostgreSQL
- golang-migrate

Install migrate:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

---

# Installation

Clone repository

```bash
git clone <repository-url>
```

Masuk ke project

```bash
cd backend1
```

Install dependency

```bash
go mod tidy
```

---

# Environment

Buat file `.env`

```env
DATABASE_URL=postgres://postgres:password@localhost:5432/backend1?sslmode=disable
```

Sesuaikan dengan database PostgreSQL milikmu.

---

# Database Migration

Migration Up

```bash
migrate -path migrations \
-database "$DATABASE_URL" up
```

Migration Down

```bash
migrate -path migrations \
-database "$DATABASE_URL" down
```

---

# Run Application

```bash
go run cmd/main.go
```

Server berjalan pada

```
http://localhost:8080
```

---

# API Endpoints

## Register

**POST**

```
/register
```

Body

```
application/x-www-form-urlencoded
```

| Field | Example |
|------|---------|
| email | admin@mail.com |
| password | admin123 |
| username | admin |
| phone | 082131213511 |

---

## Login

**POST**

```
/login
```

Body

```
application/x-www-form-urlencoded
```

| Field | Example |
|------|---------|
| email | admin@mail.com |
| password | admin123 |

---

## Get All Users

**GET**

```
/users
```

---

# Project Structure

```
.
├── cmd
│   └── main.go
├── internal
│   ├── di
│   ├── handlers
│   ├── lib
│   ├── models
│   ├── repo
│   └── service
├── migrations
├── .env
├── go.mod
├── go.sum
└── README.md
```

# API Testing Results

### Register

Berikut hasil pengujian endpoint **Register** menggunakan REST Client.

![Register Success](docs/images/register-success.png)

---

### Login

Berikut hasil pengujian endpoint **Login** menggunakan REST Client.

![Login Success](docs/images/login-success.png)

---

---

# Improvements

- JWT Authentication
- Authorization Middleware
- Refresh Token
- User Profile Endpoint
- Docker Support
- Unit Testing