# 💰 Finance Dashboard Backend

## 📌 Overview

This project is a backend system for a finance dashboard application. It provides APIs to manage users, financial records, and generate summary analytics while enforcing role-based access control.

The system follows a clean architecture (Controller → Service → Repository) to ensure scalability, maintainability, and separation of concerns.

---

## 🌐 Live Demo

**Live API:**
https://finance-backend-1zgx.onrender.com

**Swagger Documentation (Recommended for Testing):**
https://finance-backend-1zgx.onrender.com/swagger/index.html

> ⚠️ Note: The backend is deployed on Render (free tier).
> If inactive, it may take **30–60 seconds** to respond on the first request.

---

## 🔐 Test Credentials

```
Admin User:
Email: admin@example.com
```

**Steps to Test:**

1. Call `POST /api/v1/login`
2. Copy the token from response
3. Click **Authorize 🔒** in Swagger
4. Enter:

   ```
   Bearer <your_token>
   ```

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
* Ownership validation for records

### 5. Validation & Error Handling

* Input validation using Gin binding
* Meaningful error messages
* Proper HTTP status codes

### 6. Additional Features

* Rate limiting middleware
* Swagger API documentation
* Soft delete support
* Structured logging using slog

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
* No password-based authentication (email-based login)
* AutoMigrate used instead of versioned migrations
* In-memory rate limiting (not distributed)

---

## 🎯 Future Improvements

* Add password-based authentication
* Use PostgreSQL for scalability
* Implement refresh tokens
* Add caching (Redis)
* Add unit and integration tests

---

## 👨‍💻 Author

Mohit S

---

## ✅ Summary

This project demonstrates:

* Clean backend architecture
* Role-based access control (RBAC)
* Data aggregation and analytics
* Proper validation and error handling
* API documentation and deployment

It is designed to showcase backend engineering fundamentals with a practical and structured approach.
