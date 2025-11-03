# JavaDoc / Technical Documentation - GoLang User Manager

## Index

* [Overview](#overview)
* [Architecture](#architecture)
* [Modules](#modules)
* [Packages](#packages)
* [APIs and Interfaces](#apis-and-interfaces)
* [Main Flows](#main-flows)
* [Configuration and Setup](#configuration-and-setup)

---

## Overview

This project implements **user authentication and management** using **Clean Architecture** in Go.
The architecture ensures separation of concerns, testability, and maintainability through well-defined layers.

### Architectural Principles

* **Dependency Inversion**: Interfaces defined in the domain, implementations in outer layers
* **Single Responsibility**: Each module/use case has a single responsibility
* **Open/Closed**: Extensible through interfaces without modifying existing code
* **Dependency Injection**: Dependencies injected through constructors

---

## Architecture

### Clean Architecture Layers

```
┌─────────────────────────────────────────────┐
│          Presentation Layer                 │
│  (Handlers / HTTP Controllers)              │
└─────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────┐
│          Application Layer                  │
│  (Use Cases / Business Logic)               │
└─────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────┐
│          Domain Layer                       │
│  (Entities, Interfaces, Business Rules)     │
└─────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────┐
│          Infrastructure Layer               │
│  (Repository, Database, External Services)  │
└─────────────────────────────────────────────┘
```

### Dependency Directives

* **Handlers** depend on **UseCases**
* **UseCases** depend on **Domain Interfaces**
* **Repositories** implement **Domain Interfaces**
* **Domain** does not depend on anything external

---

## Modules

### 🔐 Auth Module (`internal/auth/`)

Responsible for user authentication and authorization.

#### Domain

**`internal/auth/domain/auth.go`**

* Defines interfaces for the authentication module
* Contracts for repositories and use cases

#### Repository

**`internal/auth/repository/gorm_repository.go`**

* Implements authentication data persistence using GORM
* Methods to fetch users by credentials

#### UseCase

**`internal/auth/usecase/register_user_usecase.go`**

* Creates new users
* Validates input data
* Password hashing using BCrypt

**`internal/auth/usecase/login_user_usecase.go`**

* Validates credentials
* Generates JWT tokens
* Manages sessions

**`internal/auth/usecase/refresh_token_usecase.go`**

* Renews expired JWT tokens
* Validates refresh tokens

#### Handler

**`internal/auth/handler/http.go`**

* `POST /auth/register` - User registration
* `POST /auth/login` - Login and JWT generation
* `POST /auth/refresh` - Token refresh

---

### 👥 User Module (`internal/user/`)

Full management of users, profiles, and permissions.

#### Domain

**`internal/user/domain/user.go`**

```go
type User struct {
    ID        uint
    Name      string
    Email     string
    Password  string  // BCrypt Hash
    Role      string  // user, admin, etc.
    CreatedAt time.Time
    UpdatedAt time.Time
    Profile   Profile
}

type Profile struct {
    ID     uint
    UserID uint
    Phone  string
    Avatar string
    Bio    string
}
```

**`internal/user/domain/role.go`**

* Definitions of roles and permissions
* Authentication and authorization contracts

#### Repository

**`internal/user/repository/user_repository.go`**

* User repository interface
* Methods: Create, GetByID, GetByEmail, List, Update, Delete

**`internal/user/repository/role_repository.go`**

* Role repository interface
* Methods: Assign, Revoke, ListRoles

#### UseCase - CRUD

**`internal/user/usecase/crud/create_user_usecase.go`**

* User creation with validation
* Email uniqueness check

**`internal/user/usecase/crud/get_user_by_id_usecase.go`**

* Retrieve user by ID
* Handles not found case

**`internal/user/usecase/crud/list_users_usecase.go`**

* Paginated user listing
* Optional filters

**`internal/user/usecase/crud/update_user_usecase.go`**

* User data update
* Field validation

**`internal/user/usecase/crud/delete_user_usecase.go`**

* User deletion
* Soft delete (when applicable)

#### UseCase - Authentication

**`internal/user/usecase/auth/get_profile_usecase.go`**

* Retrieves authenticated user profile

**`internal/user/usecase/auth/update_password_usecase.go`**

* Password update with validation
* Checks old password

#### UseCase - Validation

**`internal/user/usecase/validation/check_email_uniqueness_usecase.go`**

* Verifies if email already exists

**`internal/user/usecase/validation/check_user_exists_usecase.go`**

* Checks if user exists

**`internal/user/usecase/validation/validate_user_input_usecase.go`**

* Input field validation
* Data sanitization

#### UseCase - Permissions

**`internal/user/usecase/permissions/assign_role_to_user_usecase.go`**

* Assigns roles to users

**`internal/user/usecase/permissions/authorize_action_usecase.go`**

* Authorizes actions based on roles

**`internal/user/usecase/permissions/list_user_roles_usecase.go`**

* Lists user roles

**`internal/user/usecase/permissions/revoke_role_from_user_usecase.go`**

* Revokes user roles

#### UseCase - Notifications

**`internal/user/usecase/notifications/send_welcome_email_usecase.go`**

* Sends welcome email
* SMTP integration

**`internal/user/usecase/notifications/send_password_reset_usecase.go`**

* Sends password recovery link

**`internal/user/usecase/notifications/notify_user_deleted_usecase.go`**

* Notifies account deletion

**`internal/user/usecase/notifications/log_user_activity_usecase.go`**

* Logs user activity

#### Handler

**`internal/user/handler/user_handler.go`**

* `GET /users` - List all
* `GET /users/{id}` - Get by ID
* `POST /users` - Create new
* `PUT /users/{id}` - Update
* `DELETE /users/{id}` - Delete
* `PUT /users/{id}/password` - Change password

**`internal/user/handler/profile_handler.go`**

* `GET /profile` - Authenticated profile
* `PUT /profile` - Update profile

---

### 📋 Plan Module (`internal/plan/`)

Subscription plan management.

#### Domain

**`internal/plan/domain/plan.go`**

```go
type Plan struct {
    ID          uint
    Name        string
    Description string
    Price       float64
    LimitUsers  int
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

#### Repository

**`internal/plan/repository/gorm_repository.go`**

* Lists plans
* Retrieves by ID

#### UseCase

**`internal/plan/usecase/list_plans_usecase.go`**

* Lists all available plans

#### Handler

**`internal/plan/handler/http.go`**

* `GET /plans` - List plans

---

### 💳 Subscription Module (`internal/subscription/`)

User subscription management.

#### Domain

**`internal/subscription/domain/subscription.go`**

```go
type Subscription struct {
    ID        uint
    UserID    uint
    PlanID    uint
    Status    string  // active, expired, cancelled
    StartedAt time.Time
    EndsAt    time.Time
}
```

#### Repository

**`internal/subscription/repository/gorm_repository.go`**

* Creates subscriptions
* Finds by user
* Updates status

#### UseCase

**`internal/subscription/usecase/create_subscription_usecase.go`**

* Creates new subscription
* Validates limits
* Calculates dates

#### Handler

**`internal/subscription/handler/http.go`**

* `POST /subscriptions` - Create subscription

---

## Packages

### 📦 Config (`pkg/config/`)

**`pkg/config/config.go`**

* Database initialization
* Global configuration
* GORM connection with SQLite

### 🔐 JWT (`pkg/jwt/`)

**`pkg/jwt/jwt.go`**

* JWT token generation
* Token validation
* Claims extraction
* Refresh token logic

### 🔒 Hash (`pkg/hash/`)

**`pkg/hash/bcrypt.go`**

* Password hashing using BCrypt
* Hash comparison
* Cost configuration

### 📧 Email (`pkg/email/`)

**`pkg/email/smtp_service.go`**

* SMTP configuration
* Transactional email sending
* Email templates

### 📡 HTTP (`pkg/http/`)

**`pkg/http/response.go`**

* HTTP response helpers
* Standard JSON serialization
* Status codes

### 📝 Logger (`pkg/logger/`)

**`pkg/logger/logger.go`**

* Centralized logging system
* Log levels (debug, info, warn, error)
* Structured formatting

---

## APIs and Interfaces

### Domain Interfaces

#### UserRepository

```go
type UserRepository interface {
    Create(user *User) error
    GetByID(id uint) (*User, error)
    GetByEmail(email string) (*User, error)
    List() ([]User, error)
    Update(user *User) error
    Delete(id uint) error
}
```

#### PlanRepository

```go
type PlanRepository interface {
    List() ([]Plan, error)
    GetByID(id uint) (*Plan, error)
}
```

#### SubscriptionRepository

```go
type SubscriptionRepository interface {
    Create(subscription *Subscription) error
    GetByUserID(userID uint) ([]Subscription, error)
    Update(subscription *Subscription) error
}
```

### UseCase Contracts

#### RegisterUserUseCase

```go
type RegisterUserDTO struct {
    Name     string
    Email    string
    Password string
}

func (uc *RegisterUserUseCase) Execute(dto RegisterUserDTO) (*User, error)
```

#### LoginUserUseCase

```go
type LoginDTO struct {
    Email    string
    Password string
}

func (uc *LoginUserUseCase) Execute(dto LoginDTO) (string, error)
```

#### RefreshTokenUseCase

```go
func (uc *RefreshTokenUseCase) Execute(refreshToken string) (string, error)
```

---

## Main Flows

### Registration Flow

```
1. Handler receives POST /auth/register
   ↓
2. Parse body into RegisterUserDTO
   ↓
3. Call RegisterUserUseCase.Execute()
   ↓
4. UseCase validates data
   ↓
5. UseCase checks unique email (repository)
   ↓
6. UseCase hashes password (pkg/hash)
   ↓
7. UseCase creates user (repository)
   ↓
8. UseCase sends welcome email (notification)
   ↓
9. Handler returns 201 with user data
```

### Login Flow

```
1. Handler receives POST /auth/login
   ↓
2. Parse body into LoginDTO
   ↓
3. Call LoginUserUseCase.Execute()
   ↓
4. UseCase fetches user by email (repository)
   ↓
5. UseCase compares password (pkg/hash)
   ↓
6. UseCase generates JWT (pkg/jwt)
   ↓
7. Handler returns 200 with token and expires_in
```

### User Creation Flow

```
1. Handler receives POST /users
   ↓
2. Parse body into CreateUserRequest
   ↓
3. Call CreateUserUseCase.Execute()
   ↓
4. UseCase validates input (validation usecase)
   ↓
5. UseCase checks unique email (validation usecase)
   ↓
6. UseCase hashes password
   ↓
7. UseCase creates user (repository)
   ↓
8. UseCase logs activity (notification usecase)
   ↓
9. Handler returns 201 with user data
```

---

## Configuration and Setup

### Entry Point

**`cmd/app/main.go`**

* Initializes Fiber app
* Configures database
* Runs migrations
* Injects dependencies
* Registers routes
* Starts server on port 8080

### Dependency Bootstrap

```
main() {
    app := fiber.New()
    
    // Database
    config.InitDatabase()
    config.DB.AutoMigrate(&User{}, &Profile{}, &Plan{}, &Subscription{})
    
    // Repositories
    userRepo := ur.NewGormUserRepository(config.DB)
    planRepo := pr.NewGormPlanRepository(config.DB)
    subRepo := sr.NewGormSubscriptionRepository(config.DB)
    
    // UseCases
    listUsers := ucrud.NewListUsersUseCase(userRepo)
    createUser := ucrud.NewCreateUserUseCase(userRepo)
    // ... other use cases
    
    // Handlers
    userHandler := uh.NewUserHandler(listUsers, createUser, ...)
    authHandler := authh.NewAuthHandler(registerUC, loginUC, refreshUC)
    
    // Routes
    api := app.Group("/")
    userHandler.RegisterRoutes(api)
    authHandler.RegisterRoutes(api)
    
    // Docs
    docs.ServeDocs(app)
    
    app.Listen(":8080")
}
```

### Documentation

**`docs/`**

* `generator.go` - Serves documentation files
* `openapi.yaml` - OpenAPI 3.1 specification
* `scalar/` - Scalar UI for visualization

Access: `http://localhost:8080/docs`

---

## Code Conventions

### Naming

* **Packages**: lowercase, single word
* **Types**: PascalCase (User, Plan, Subscription)
* **Interfaces**: PascalCase with appropriate suffix (UserRepository, UseCase)
* **Functions**: PascalCase for public, camelCase for private
* **Variables**: camelCase

### File Structure

* One file per entity/interface in the domain
* One file per use case
* One main file per handler
* Repositories grouped by domain

### Error Handling

* UseCases return descriptive errors
* Handlers convert errors into HTTP status codes
* No panics; always return errors

### Validation

* Input validation within UseCases
* Dedicated validation use cases when needed
* Clear and specific error messages

---

## Future Improvements

### Security

* [ ] Rate limiting
* [ ] Configurable CORS
* [ ] Helmet headers
* [ ] CSRF protection
* [ ] Stronger input validation

### Features

* [ ] JWT authentication middleware
* [ ] Real pagination in listings
* [ ] Advanced filters and search
* [ ] Soft delete
* [ ] Audit logs

### Infrastructure

* [ ] Structured migrations
* [ ] Data seeding
* [ ] Docker Compose
* [ ] Environment-based configuration
* [ ] CI/CD

### Tests

* [ ] Unit tests for UseCases
* [ ] Integration tests for Handlers
* [ ] Repository tests with mock DB
* [ ] Test coverage > 80%

---

**Last updated**: November 2025
