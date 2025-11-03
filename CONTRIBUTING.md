# Contribution Guide

Thank you for considering contributing to the **GoLang User Manager** project!
This is an educational and training project, and all contributions are welcome.

## 📋 Table of Contents

* [Code of Conduct](#code-of-conduct)
* [How to Contribute](#how-to-contribute)
* [Code Standards](#code-standards)
* [Commit Structure](#commit-structure)
* [Pull Request Process](#pull-request-process)

---

## Code of Conduct

### Our Commitments

* Be respectful and professional
* Accept constructive criticism gracefully
* Focus on what’s best for the community
* Show empathy towards other community members

---

## How to Contribute

### Reporting Issues

If you find a bug or have a suggestion:

1. Check if the issue has already been reported in the [Issues](../../issues)
2. Create a new issue with:

    * **Clear and descriptive title**
    * **Detailed description** of the problem
    * **Steps to reproduce** (if it’s a bug)
    * **Expected vs. actual behavior**
    * **Environment details** (Go version, OS, etc.)

### Suggesting Improvements

To suggest new features:

1. Check if there’s already a related issue
2. Open a new issue labeled `enhancement`
3. Explain:

    * **The problem it solves**
    * **The proposed solution**
    * **The benefits** for the project

### Contributing Code

#### Development Environment Setup

1. Fork the repository
2. Clone your fork:

```bash
git clone https://github.com/tarikvillalobos/golang-user-manager.git
cd golang-user-manager
```

3. Create a branch:

```bash
git checkout -b feature/your-feature
# or
git checkout -b fix/your-bugfix
```

4. Make your changes

5. Test your changes:

```bash
go run cmd/app/main.go
```

6. Commit following our conventions

7. Push to your fork:

```bash
git push origin feature/your-feature
```

8. Open a Pull Request

---

## Code Standards

### Clean Architecture

* ✅ Maintain separation between layers
* ✅ UseCases should not depend on HTTP details
* ✅ Repositories must implement domain interfaces
* ✅ Domain should not depend on implementations

### Naming

* **Packages**: lowercase, single word
* **Types**: PascalCase
* **Public functions**: PascalCase
* **Private functions**: camelCase
* **Variables**: camelCase
* **Constants**: PascalCase or UPPER_CASE

### File Structure

```go
// Domain
package domain

type User struct { ... }
type UserRepository interface { ... }

// Repository
package repository

type GormUserRepository struct { ... }
func NewGormUserRepository(db *gorm.DB) *GormUserRepository { ... }

// UseCase
package usecase

type CreateUserUseCase struct { ... }
func NewCreateUserUseCase(repo UserRepository) *CreateUserUseCase { ... }
func (uc *CreateUserUseCase) Execute(dto CreateUserDTO) (*User, error) { ... }

// Handler
package handler

type UserHandler struct { ... }
func NewUserHandler(uc UseCase) *UserHandler { ... }
func (h *UserHandler) RegisterRoutes(r fiber.Router) { ... }
```

### Go Specific

* Use `go fmt` for formatting
* Use `go vet` to detect common issues
* Avoid unused imports
* Comment public functions
* Use `context.Context` in async operations
* Return explicit errors, never use panics

### Error Handling

```go
// ✅ Good
func (uc *CreateUserUseCase) Execute(dto DTO) (*User, error) {
    user, err := uc.repo.Create(dto)
    if err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }
    return user, nil
}

// ❌ Bad
func (uc *CreateUserUseCase) Execute(dto DTO) (*User, error) {
    user := uc.repo.Create(dto) // Returns error?
    return user, nil
}
```

### Validation

```go
// ✅ Proper validation
func (uc *CreateUserUseCase) Execute(dto DTO) (*User, error) {
    if dto.Email == "" {
        return nil, errors.New("email is required")
    }
    if !isValidEmail(dto.Email) {
        return nil, errors.New("invalid email format")
    }
    // ...
}
```

---

## Commit Structure

We follow the [Conventional Commits](https://www.conventionalcommits.org/) standard:

```
<type>[optional scope]: <description>

[optional body]

[optional footer]
```

### Types

* `feat`: New feature
* `fix`: Bug fix
* `docs`: Documentation changes
* `style`: Formatting, whitespace, semicolons, etc.
* `refactor`: Code refactoring
* `test`: Adding or fixing tests
* `chore`: Build, configuration, or maintenance changes

### Scopes (Optional)

* `auth`: Authentication module
* `user`: User module
* `plan`: Plan module
* `subscription`: Subscription module
* `config`: Configuration
* `docs`: Documentation
* `handler`: HTTP controllers

### Examples

```bash
feat(auth): add automatic refresh token
fix(user): fix duplicate email validation
docs: update README with setup instructions
refactor(handler): simplify HTTP response logic
test(user): add unit tests for CreateUserUseCase
chore: update Go dependencies
```

### Good Commit Messages

* ✅ `feat(auth): implement refresh token`
* ✅ `fix(user): fix bug in email search`
* ✅ `docs: add API usage examples`

### Bad Commit Messages

* ❌ `fix stuff`
* ❌ `updates`
* ❌ `a`
* ❌ `FIX BUG!!!!`

---

## Pull Request Process

### Before Opening a PR

1. ✅ Code follows established standards
2. ✅ Commits follow Conventional Commits
3. ✅ Feature has been manually tested
4. ✅ No linting/formatting errors
5. ✅ README updated (if necessary)

### Creating the PR

1. **Descriptive Title**

   ```
   feat(auth): implement JWT authentication middleware
   ```

2. **Clear Description**

    * What was implemented
    * Why it was implemented
    * How it works
    * How to test it

3. **PR Template**

```markdown
## Description
Brief description of what was implemented.

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation

## How It Was Tested
Describe how you tested the change:
- [ ] Manual test on endpoint X
- [ ] Integration test
- [ ] Edge case verification

## Checklist
- [ ] Code follows project standards
- [ ] Self-review completed
- [ ] Comments added for complex logic
- [ ] Documentation updated
- [ ] No build warnings/errors
- [ ] Commits follow Conventional Commits
```

### Review Process

1. Maintainers will review the PR
2. There may be requests for changes
3. Respond to comments and feedback
4. Make necessary updates
5. The PR will be merged once approved

---

## Areas That Need Contributions

### High Priority

* [ ] JWT authentication middleware
* [ ] Unit tests for UseCases
* [ ] Integration tests for Handlers
* [ ] Stronger input validation

### Medium Priority

* [ ] Rate limiting
* [ ] Real pagination
* [ ] Soft delete
* [ ] Audit logs
* [ ] Structured migrations

### Low Priority

* [ ] Docker Compose
* [ ] CI/CD
* [ ] Performance improvements
* [ ] Load testing

---

## Useful Resources

### Documentation

* [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
* [Go Code Review](https://github.com/golang/go/wiki/CodeReviewComments)
* [Conventional Commits](https://www.conventionalcommits.org/)
* [Fiber Documentation](https://docs.gofiber.io/)

### Tools

```bash
# Formatting
go fmt ./...

# Linting
golangci-lint run

# Tests
go test ./...

# Build
go build -o bin/app cmd/app/main.go
```

---

## Questions?

If you have any questions about contributing:

* Open an issue labeled `question`
* Refer to the documentation in `README.md` and `JAVADOC.md`

---

**Thank you for contributing! 🎉**
