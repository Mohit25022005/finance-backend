# 💰 Finance Dashboard Backend

## 📌 Overview

This project is a backend system for a finance dashboard application. It provides APIs to manage users, financial records, and generate summary analytics while enforcing role-based access control.

The system is designed with a clean architecture (Controller → Service → Repository) to ensure scalability, maintainability, and clear separation of concerns.

---

## 🚀 Features

### 1. User & Role Management

* Create, update, and delete users
* Assign roles: **Admin, Analyst, Viewer**
* Enable/disable users (active/inactive)
* Role-based access control (RBAC)

### 2. Financial Records Management

* Create, update, delete records
* Fields:

  * Amount
  * Type (Income / Expense)
  * Category
  * Date
  * Notes
* Filtering:

  * By type
  * By category
  * Search support
  * Pagination

### 3. Dashboard Analytics

* Total income
* Total expenses
* Net balance
* Category-wise totals
* Recent activity
* Monthly trends

### 4. Authentication & Security

* JWT-based authentication
* Protected routes
* Role-based authorization middleware
* Ownership checks for records

### 5. Validation & Error Handling

* Input validation using Gin binding
* Meaningful error messages
* Proper HTTP status codes

### 6. Additional Features

* Rate limiting middleware
* Swagger API documentation
* Soft delete support
* Logging with slog

---

## 🏗️ Tech Stack

* **Language:** Go (Golang)
* **Framework:** Gin
* **ORM:** GORM
* **Database:** SQLite
* **Authentication:** JWT
* **Documentation:** Swagger

---

## 📂 Project Structure

```
.
├── controllers     # HTTP handlers
├── services        # Business logic
├── repository      # Database queries
├── models          # Data models & DTOs
├── middleware      # Auth, RBAC, rate limiting
├── routes          # Route definitions
├── config          # DB connection
├── utils           # JWT utilities
```

---

## 🔐 Roles & Permissions

| Role    | Permissions                   |
| ------- | ----------------------------- |
| Viewer  | View dashboard only           |
| Analyst | View records + dashboard      |
| Admin   | Full access (users + records) |

---

## 🔑 Authentication

* Login using email
* Receive JWT token
* Pass token in headers:

```
Authorization: Bearer <token>
```

---

## 📡 API Endpoints

### Auth

* `POST /api/v1/login`

### Records

* `POST /api/v1/records`
* `GET /api/v1/records`
* `PUT /api/v1/records/:id`
* `DELETE /api/v1/records/:id`

### Users

* `POST /api/v1/users`
* `GET /api/v1/users`
* `GET /api/v1/users/:id`
* `PUT /api/v1/users/:id`
* `DELETE /api/v1/users/:id`

### Dashboard

* `GET /api/v1/dashboard`

---

## ⚙️ Setup Instructions

### 1. Clone repository

```
git clone <your-repo-url>
cd finance-backend
```

### 2. Install dependencies

```
go mod tidy
```

### 3. Run the server

```
go run main.go
```

### 4. Access Swagger

```
http://localhost:8084/swagger/index.html
```

---

## 🗄️ Database

* Uses **SQLite** for simplicity
* Auto-migrations handled via GORM
* Suitable for development and testing

---

## ⚠️ Assumptions & Trade-offs

* SQLite used instead of PostgreSQL for simplicity
* No password authentication (email-based login only)
* AutoMigrate used instead of versioned migrations
* Basic rate limiting (in-memory)

---

## 🎯 Future Improvements

* Add password-based authentication
* Use PostgreSQL for scalability
* Implement refresh tokens
* Add caching (Redis)
* Write unit & integration tests

---

## 👨‍💻 Author

Mohit S

---

## ✅ Summary

This project demonstrates:

* Clean backend architecture
* Role-based access control
* Data aggregation & analytics
* Proper validation and error handling

It is designed to showcase backend engineering fundamentals in a practical and structured way.
