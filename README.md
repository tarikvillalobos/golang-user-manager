# GoLang User Manager - User Management with Clean Architecture

[![Go Version](https://img.shields.io/badge/Go-1.22-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## 📋 Description

User authentication and management project developed in Go, using **Clean Architecture** as the architectural pattern.
This is a study and training project for backend implementations, focused on:

* **Complete JWT Authentication** (register, login, refresh token)
* **Clean Architecture** with clear separation of responsibilities
* **API Documentation** using Scalar for endpoint visualization and testing
* **Modular and scalable architecture**

## 🏗️ Architecture

The project follows Clean Architecture principles, organizing code into well-defined layers:

```
golang-user-manager/
├── cmd/app/           # Main application and dependency setup
├── internal/          # Business logic (not exposed externally)
│   ├── auth/          # Authentication module
│   ├── user/          # User module
│   ├── profile/       # Profile module
│   ├── plan/          # Plan module
│   └── subscription/  # Subscription module
├── pkg/               # Shared packages and utilities
└── docs/              # API documentation (Scalar/OpenAPI)
```

### Clean Architecture Layers

Each internal module follows this structure:

```
module/
├── domain/       # Domain entities and interfaces
├── repository/   # Data persistence implementation
├── usecase/      # Business logic and use cases
└── handler/      # HTTP controllers and adapters
```

## 🚀 Technologies

* **[Go Fiber](https://github.com/gofiber/fiber)** – Fast and expressive web framework
* **[GORM](https://gorm.io/)** – ORM for Go
* **[SQLite](https://www.sqlite.org/)** – Embedded database
* **[JWT](https://github.com/golang-jwt/jwt)** – Authentication with JSON Web Tokens
* **[BCrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)** – Password hashing
* **[Scalar](https://github.com/scalar/scalar)** – Interactive API documentation
* **[OpenAPI 3.1](https://swagger.io/specification/)** – API specification standard

## 📦 Installation

### Prerequisites

* Go 1.22 or higher
* Git

### Setup

1. Clone the repository:

```bash
git clone <repository-url>
cd golang-user-manager
```

2. Install dependencies:

```bash
go mod download
```

3. Run the project:

```bash
go run cmd/app/main.go
```

4. Access the interactive documentation:

```
http://localhost:8080/docs
```

## 📚 Features

### 🔐 Authentication

* **User Registration**: Create new accounts with validation
* **Login**: Authenticate with email and password
* **Refresh Token**: Renew JWT tokens
* **Password Hashing**: Secure storage using BCrypt

### 👥 User Management

* **Full CRUD**: Create, list, fetch, update, and delete users
* **Profiles**: Manage extended user information
* **Password Change**: Secure password update
* **Roles**: Permissions and role-based access system

### 📋 Plans and Subscriptions

* **Plans**: List available service plans
* **Subscriptions**: Create and manage user subscriptions

## 🌐 Endpoints

| Method   | Endpoint               | Description               |
| -------- | ---------------------- | ------------------------- |
| `POST`   | `/auth/register`       | Register a new user       |
| `POST`   | `/auth/login`          | User login                |
| `POST`   | `/auth/refresh`        | Refresh JWT token         |
| `GET`    | `/users`               | List all users            |
| `GET`    | `/users/{id}`          | Get user by ID            |
| `POST`   | `/users`               | Create a new user         |
| `PUT`    | `/users/{id}`          | Update user               |
| `DELETE` | `/users/{id}`          | Delete user               |
| `PUT`    | `/users/{id}/password` | Update password           |
| `GET`    | `/profile`             | Get authenticated profile |
| `PUT`    | `/profile`             | Update profile            |
| `GET`    | `/plans`               | List plans                |
| `POST`   | `/subscriptions`       | Create subscription       |

For full documentation and examples, visit: `http://localhost:8080/docs`

## 🗂️ Project Structure

### Main Modules

#### **Auth** (`internal/auth/`)

Handles authentication and authorization:

* **Domain**: Authentication interfaces and contracts
* **Handler**: Endpoints `/auth/*`
* **Repository**: Authentication data persistence
* **Usecase**: Register, login, and refresh token logic

#### **User** (`internal/user/`)

Full user and profile management:

* **Domain**: User, Profile, and Role entities
* **Handler**: Endpoints `/users/*` and `/profile`
* **Repository**: User and profile CRUD operations
* **Usecase**:

    * CRUD (create, read, update, delete)
    * Authentication (get profile, update password)
    * Validation (check email, validate input)
    * Permissions (assign role, authorize, revoke)
    * Notifications (welcome email, password reset, etc.)

#### **Plan** (`internal/plan/`)

Service plan management:

* **Domain**: Plan entity
* **Handler**: Endpoints `/plans/*`
* **Repository**: Plan persistence
* **Usecase**: Plan listing

#### **Subscription** (`internal/subscription/`)

Subscription management:

* **Domain**: Subscription entity
* **Handler**: Endpoints `/subscriptions/*`
* **Repository**: Subscription persistence
* **Usecase**: Subscription creation

### Shared Packages (`pkg/`)

* **config**: Database configuration
* **hash**: Password hashing utilities (BCrypt)
* **jwt**: JWT token generation and validation
* **http**: HTTP response helpers
* **logger**: Logging system
* **email**: SMTP service for notifications

## 🔧 Configuration

### Database

The project uses SQLite by default.
An `app.db` file is automatically created at the project root.
To change the database, edit `pkg/config/config.go`.

### Environment Variables

Currently, configuration values are hardcoded.
For production, consider externalizing:

* JWT secret keys
* SMTP credentials
* Database URLs

## 🧪 Testing

To run tests (when implemented):

```bash
go test ./...
```

## 📖 Additional Documentation

* [Scalar UI](http://localhost:8080/docs) – Interactive API documentation
* [OpenAPI Spec](http://localhost:8080/openapi.yaml) – Full API specification
* [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) – Architectural principles

## 🤝 Contributing

This is a study and training project. Contributions are welcome!

1. Fork the project
2. Create a feature branch (`git checkout -b feature/MyFeature`)
3. Commit your changes (`git commit -m 'feat: Add MyFeature'`)
4. Push to the branch (`git push origin feature/MyFeature`)
5. Open a Pull Request

## 📝 License

This project is under the MIT license. See the LICENSE file for details.

## 👨‍💻 Author

Developed as a training project in Go and Clean Architecture.

---

**Note**: This is an educational project for backend implementation practice.
It is not recommended for production use without proper security and scalability adjustments.
