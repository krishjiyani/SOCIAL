# Social App API – Golang (In Progress)

A scalable and modular RESTful API built in Go, designed for a minimal social networking platform. This backend system handles core social media features such as user profiles, post interactions, and secure request handling. The project is focused on clean architecture, robust validation, and high performance.

## 🚧 Project Status
This project is currently under active development. Core post-related functionalities are implemented. Upcoming features include authentication, post interactions (likes/comments), caching, and testing.

---

## 🔧 Tech Stack & Libraries

- **Language**: Go (Golang)
- **Framework**: `net/http` with [Chi Router](https://github.com/go-chi/chi)
- **Database**: PostgreSQL (abstracted behind a store interface)
- **Validation**: [`go-playground/validator`](https://github.com/go-playground/validator)
- **Middleware**: Logging, Real IP, Panic Recovery, Timeout, Context Injection
- **Architecture**: REST principles, clean routing, modular handlers

---

## 📁 Features Implemented

### ✅ Core Routing
- `GET /health` – Health check endpoint
- `POST /posts` – Create new post
- `GET /posts/{postID}` – Fetch single post
- `PATCH /posts/{postID}` – Update post
- `DELETE /posts/{postID}` – Delete post
- `GET /users/{userID}` – Retrieve user profile

### ✅ Middleware & Server Setup
- Context-aware middleware stack
- Logging and panic recovery
- Custom context injection for preloading post data

### ✅ Input Validation
- Strong payload validation using `validator.New()` for all endpoints
- Structured error handling with consistent JSON responses

### ✅ Config & Environment
- Config management using `.env` file for DB connection and secrets
- Modular environment loading

---

## 🧪 Upcoming Features

- 🔐 User Authentication & Authorization (JWT-based)
- ❤️ Post Interactions (likes, comments)
- 🧠 Caching layer for improved performance
- 🧪 Full Unit & Integration Test Coverage
- 📦 Swagger/OpenAPI documentation

---

## 🚀 Getting Started

### Prerequisites
- Go 1.20+
- PostgreSQL running locally or remotely
- `.env` file with environment variables
