# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/).

## [Unreleased]

### Added

* Complete project documentation (README.md)
* Detailed technical JavaDoc (JAVADOC.md)
* Contribution guide (CONTRIBUTING.md)
* Changelog (CHANGELOG.md)
* MIT License (LICENSE)

### Improved

* Go code formatting across the entire project

## [0.1.0] - 2025-11-02

### Added

* Initial project structure with Clean Architecture
* Authentication Module (Auth)

    * User registration
    * JWT login
    * Refresh token
* User Module (User)

    * Full user CRUD
    * Profile management
    * Password change
    * Validations and notifications
* Plan Module (Plan)

    * Plan listing
* Subscription Module (Subscription)

    * Subscription creation
* Shared packages (pkg)

    * Database configuration
    * BCrypt hashing
    * JWT
    * HTTP helpers
    * Logger
    * SMTP email
* Documentation with Scalar

    * OpenAPI 3.1 specification
    * Interactive testing interface
* Integration with Fiber framework
* SQLite database with GORM
* Automatic migrations

### Notes

* Initial project release
* Educational and training project
* Focused on implementing Clean Architecture in Go
* Foundation for future improvements and features

---

## Types of Changes

* **Added**: for new features
* **Changed**: for changes in existing functionality
* **Deprecated**: for features soon to be removed
* **Removed**: for removed features
* **Fixed**: for bug fixes
* **Security**: for security vulnerabilities

---

## Version Format

Versions follow the [MAJOR.MINOR.PATCH] pattern:

* **MAJOR**: Incompatible API changes
* **MINOR**: Backward-compatible new functionality
* **PATCH**: Backward-compatible bug fixes

---

## Useful Links

* [README.md](./README.md) - Main documentation
* [JAVADOC.md](./JAVADOC.md) - Technical documentation
* [CONTRIBUTING.md](./CONTRIBUTING.md) - Contribution guide
* [Issues](../../issues) - Report problems
* [Pull Requests](../../pulls) - Contribute code
