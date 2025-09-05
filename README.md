<img width="1919" height="973" alt="image" src="https://github.com/user-attachments/assets/ad7ed89d-5dd2-46ba-8d06-acb2f58add1f" />

# Snippetbox

A web application for creating and sharing code snippets built with Go and MySQL.

## Features

- **Create Snippets**: Users can create code snippets with titles, content, and expiration dates
- **View Snippets**: Browse and view individual snippets
- **User Authentication**: User registration, login, and logout functionality
- **Session Management**: Secure session handling with MySQL storage
- **Flash Messages**: User feedback via flash messages
- **HTTPS Support**: TLS encryption for secure connections
- **Template Caching**: Efficient template rendering with caching
- **Middleware**: Custom middleware for logging, panic recovery, and security headers
- **Form Validation**: Client and server-side form validation
- **CSRF Protection**: Cross-site request forgery protection

## Tech Stack

- **Backend**: Go 1.22
- **Database**: MySQL
- **Frontend**: HTML templates with CSS and JavaScript
- **Session Store**: MySQL-backed sessions
- **Router**: httprouter
- **Middleware**: Alice for middleware chaining
- **Security**: nosurf for CSRF protection, SCS for session management

## Project Structure

```
snippetbox/
├── cmd/web/           # Web application entry point
│   ├── handlers.go    # HTTP handlers
│   ├── helpers.go     # Helper functions
│   ├── main.go        # Application entry point
│   ├── middleware.go  # Custom middleware
│   ├── routes.go      # Route definitions
│   └── templates.go   # Template management
├── internal/
│   ├── models/        # Data models and database logic
│   │   ├── mocks/     # Mock implementations for testing
│   │   ├── testdata/  # Test data and SQL scripts
│   │   ├── errors.go  # Custom error types
│   │   ├── snippets.go # Snippet model
│   │   └── users.go   # User model
│   └── validator/     # Form validation logic
├── ui/
│   ├── html/          # HTML templates
│   │   ├── base.html  # Base template
│   │   ├── pages/     # Page templates
│   │   └── partials/  # Partial templates
│   ├── static/        # Static assets (CSS, JS, images)
│   └── efs.go         # Embedded filesystem
├── tls/               # TLS certificates
└── go.mod             # Go module definition
```

## Prerequisites

- Go 1.22 or higher
- MySQL 8.0 or higher
- Git

## Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/ASHUTOSH-SWAIN-GIT/snippetbox.git
   cd snippetbox
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Set up MySQL database**

   ```sql
   CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
   CREATE USER 'web'@'localhost' IDENTIFIED BY 'pass';
   GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'localhost';
   ```

4. **Create database tables**
   ```bash
   mysql -u web -p snippetbox < internal/models/testdata/setup.sql
   ```

## Usage

### Running the Application

**Development (HTTP)**

```bash
go run ./cmd/web
```

**Production (HTTPS)**

```bash
go run ./cmd/web -addr=":443" -dsn="web:your_password@/snippetbox?parseTime=true"
```

The application will be available at:

- HTTP: http://localhost:4000
- HTTPS: https://localhost:4000

### Command Line Flags

- `-addr`: Network address (default: ":4000")
- `-dsn`: MySQL data source name (default: "web:pass@/snippetbox?parseTime=true")

### API Endpoints

- `GET /` - Home page with recent snippets
- `GET /snippet/view/:id` - View a specific snippet
- `GET /snippet/create` - Create snippet form
- `POST /snippet/create` - Submit new snippet
- `GET /user/signup` - User registration form
- `POST /user/signup` - Submit registration
- `GET /user/login` - User login form
- `POST /user/login` - Submit login
- `POST /user/logout` - User logout
- `GET /ping` - Health check endpoint

## Testing

**Run all tests**

```bash
go test ./...
```

**Run tests with verbose output**

```bash
go test -v ./...
```

**Run specific package tests**

```bash
go test -v ./cmd/web/
go test -v ./internal/models/
```

## Configuration

### Database Configuration

The application uses MySQL with the following default settings:

- Host: localhost
- Database: snippetbox
- Username: web
- Password: pass

### Session Configuration

- Session lifetime: 12 hours
- Session store: MySQL
- Cookie settings: HttpOnly, SameSite=Lax

### Security Features

- CSRF protection on all state-changing requests
- Secure headers (X-Frame-Options, X-XSS-Protection, etc.)
- HTTPS redirect in production
- Session-based authentication
- Password hashing with bcrypt

## Development

### Adding New Features

1. Create new handlers in `cmd/web/handlers.go`
2. Add routes in `cmd/web/routes.go`
3. Create templates in `ui/html/pages/`
4. Add models in `internal/models/`
5. Write tests for new functionality

### Database Migrations

To add new database changes:

1. Create migration SQL files
2. Update the setup.sql for new installations
3. Document changes in the README

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built following the "Let's Go" book by Alex Edwards
- Uses the excellent httprouter for routing
- Session management powered by SCS
- CSRF protection via nosurf

## Contact

- GitHub: [@ASHUTOSH-SWAIN-GIT](https://github.com/ASHUTOSH-SWAIN-GIT)
- Project Link: [https://github.com/ASHUTOSH-SWAIN-GIT/snippetbox](https://github.com/ASHUTOSH-SWAIN-GIT/snippetbox)
