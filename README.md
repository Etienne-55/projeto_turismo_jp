# Tourism Platform API

A REST API for managing tourist trips built with Go and Gin. Features JWT authentication, role-based access control, and comprehensive testing. Designed for learning backend architecture patterns.

## Features

* **JWT Authentication** - Secure token-based auth with role support
* **Role-based Access** - Admin and user roles with protected endpoints
* **CRUD Operations** - Complete trip management system
* **Swagger Documentation** - Interactive API docs at `/swagger`
* **Dependency Injection** - Clean architecture with repository pattern
* **Comprehensive Testing** - 80%+ coverage with mocks
* **SQLite Database** - Lightweight, file-based persistence

## Installation
```bash
git clone https://github.com/Etienne-55/projeto-turismo.git
cd projeto-turismo/backend
go mod download
```

## Usage
```bash
# Run the server
go run main.go

# Server starts on http://localhost:8080
# Swagger docs at http://localhost:8080/swagger/index.html
```

## API Endpoints

### Authentication
* `POST /signup` - Create new user account
* `POST /login` - Login and receive JWT token

### Trips (Authenticated)
* `POST /trip` - Create a new trip
* `GET /trips` - Get user's trips
* `GET /trip/:id` - Get specific trip
* `PUT /trip/:id` - Update trip
* `DELETE /trip/:id` - Delete trip

### Admin (Admin Role Required)
* `GET /admin/all-trips` - Get all trips from all users
* `GET /admin/users` - Get all users

## Authentication

Include JWT token in requests:
```bash
curl -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  http://localhost:8080/trips
```

## Project Structure
```
backend/
├── controllers/     # HTTP handlers
├── models/         # Data structures
├── repositories/   # Database layer
├── middleware/     # Auth & admin middleware
├── routes/         # Route definitions
├── utils/          # JWT & password hashing
├── db/             # Database initialization
└── docs/           # Swagger generated docs
```

## Testing
```bash
# Run all tests
go test ./...

# Run with coverage
go test ./... -cover

# Run specific package
go test ./controllers -v
```

## Technologies

* **Gin** - HTTP web framework
* **JWT-Go** - JSON Web Token authentication
* **Swagger** - API documentation
* **SQLite** - Database
* **Testify** - Testing assertions

## Development
```bash
# Generate Swagger docs after changes
swag init

# Run tests
go test ./...

# Build
go build -o tourism-api
```

## Admin Setup

Default admin account is created on first run:
* Email: `admin@proton.me`
* Password: `admin123`

**Change this immediately in production!**

## Configuration

Database file: `api.db` (created automatically)

JWT secret: Configured in `utils/jwt.go` (use environment variables in production)

## License

MIT
