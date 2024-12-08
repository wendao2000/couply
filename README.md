# Couply - Dating App Backend

A Go-based dating app backend service built with Clean Architecture principles. This implementation focuses on authentication services, with additional features planned for future development.

## Project Structure

This service follows Clean Architecture patterns for clear separation of concerns:

### Directory Layout

```text
internal/
  ├── config/       # Configuration and environment setup
  ├── constants/    # Application constants and test helpers
  ├── database/     # Database connection and management
  ├── handlers/     # HTTP request handlers (auth, profile)
  ├── models/       # Data models and entities
  ├── redis/        # Redis client setup
  ├── repositories/ # Data access layer
  ├── response/     # Standardized API responses
  └── services/     # Business logic layer

pkg/
  ├── errs/         # Error handling utilities
  └── utils/        # Common utilities
```

## Running the Service

### Prerequisites

- Go 1.21 or higher

### Steps to Run

1. Clone the repository

    ```bash
    git clone [repository-url]
    cd couply
    ```

2. Start the service

    ```bash
    go run main.go
    ```

The server will start at `http://localhost:2025`.

### Available Endpoints

- Sign Up: `POST http://localhost:2025/signup`
- Sign In: `POST http://localhost:2025/signin`

## Dependencies

- **chi/v5**: HTTP routing
- **jwt/v5**: Authentication
- **godotenv**: Environment configuration
- **gorm**: Database ORM
- **go-redis/v9**: Caching (planned)

The service uses SQLite for development simplicity, requiring no additional database setup.
