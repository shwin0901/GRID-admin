# GRID Admin

## Tech Stack

- **Web Framework**: [Gin](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Token)
- **Password Encryption**: bcrypt

## Project Structure

```
go-admin/
├── main.go                 # Entry point
├── go.mod                  # Go module definition
├── config/
│   └── config.go           # Configuration management
├── database/
│   └── db.go               # Database connection
├── models/
│   └── user.go             # User model
├── handlers/
│   └── auth.go             # Authentication handlers
├── middleware/
│   └── jwt.go              # JWT middleware
├── routes/
│   └── routes.go           # Route definitions
└── utils/
    └── response.go         # Unified response format
```

## Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL 12+

### Database Setup

Create a database named `go_admin`:

```sql
CREATE DATABASE go_admin;
```

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| DB_HOST | Database host | localhost |
| DB_PORT | Database port | 5432 |
| DB_USER | Database user | postgres |
| DB_PASSWORD | Database password | postgres |
| DB_NAME | Database name | go_admin |
| JWT_SECRET | JWT signing key | your-secret-key |
| SERVER_PORT | Server port | 8080 |

### Run the Server

```bash
cd go-admin
go run main.go
```

The server will start at `http://localhost:8080`

## API Documentation

### Public Endpoints (No Authentication Required)

#### Register

```
POST /api/auth/register
```

**Request Body:**
```json
{
  "username": "admin",
  "password": "123456",
  "email": "admin@example.com"
}
```

**Response:**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "username": "admin",
    "email": "admin@example.com",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### Login

```
POST /api/auth/login
```

**Request Body:**
```json
{
  "username": "admin",
  "password": "123456"
}
```

**Response:**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user_info": {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com"
    }
  }
}
```

### Protected Endpoints (JWT Required)

#### Get User Profile

```
GET /api/user/profile
Authorization: Bearer <token>
```

**Response:**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "username": "admin",
    "email": "admin@example.com",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

## Quick Test with cURL

```bash
# Register a new user
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456","email":"admin@example.com"}'

# Login
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456"}'

# Get user profile (replace <token> with the token from login response)
curl http://localhost:8080/api/user/profile \
  -H "Authorization: Bearer <token>"
```

## License

MIT
