# Gin REST API with JWT Authentication

A robust REST API built with Go (Gin framework) featuring JWT authentication, PostgreSQL integration, and a clean architecture pattern.

## Features

- 🚀 Built with [Gin Web Framework](https://github.com/gin-gonic/gin)
- 🔐 JWT Authentication
- 📦 PostgreSQL with GORM
- 🏗️ Clean Architecture
- ⚙️ Environment Configuration
- 🔒 Password Hashing
- 📝 Structured Logging

## Project Structure

```
my-gin-project/
├── config/             # Configuration settings
│   ├── config.go      # App configuration
│   └── database.go    # Database configuration
├── controllers/        # Request handlers
│   ├── auth_controller.go
│   └── user_controller.go
├── models/            # Data models
│   └── user.go
├── routes/            # Route definitions
│   ├── auth_routes.go
│   └── user_routes.go
├── services/          # Business logic
│   ├── auth_service.go
│   └── user_service.go
├── db/                # Database setup
│   └── db.go
├── .env               # Environment variables
├── .env.example       # Example environment variables
└── main.go           # Application entry point
```

## Prerequisites

- Go 1.16 or higher
- PostgreSQL
- Git

## Installation

1. Clone the repository
```bash
git clone https://github.com/RuberDucky/GoRestAPI.git
cd GoRestAPI
```

2. Install dependencies
```bash
go mod download
```

3. Set up environment variables
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Run the application
```bash
go run main.go
```

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
# Server Configuration
SERVER_PORT=8080
GIN_MODE=debug

# Database Configuration
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=your_database
DB_PORT=5432

# JWT Configuration
JWT_SECRET_KEY=your-secret-key
```

## API Endpoints

### Authentication
- `POST /auth/login` - User login
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```

### Users
- `POST /users` - Create new user
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }
  ```
- `GET /users` - Get all users

## Database Schema

### Users Table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    age INT DEFAULT 0,
    is_active BOOLEAN DEFAULT true,
    last_login TIMESTAMP
);
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [JWT Go](https://github.com/golang-jwt/jwt)

## Contact

Your Name - [@Zain](zaindev@duck.com)
