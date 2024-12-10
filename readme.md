# Golang Travel Booking API

This is a Go-based web application for a travel booking system. It provides a RESTful API for managing user authentication, user profiles, and other travel-related functionality.

## Table of Contents

- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Application](#running-the-application)
  - [Running Migrations](#running-migrations)
- [API Documentation](#api-documentation)
  - [Authentication](#authentication)
  - [User Management](#user-management)
- [Testing](#testing)
- [Deployment](#deployment)
- [Built With](#built-with)
- [Contributing](#contributing)
- [License](#license)

## Project Structure

The project follows a clean architecture approach with the following directories:

- `cmd`: Contains the main entry points for the application.
  - `api`: The main web application that exposes the API endpoints.
  - `migration`: A command-line tool for running database migrations.
- `database`: Contains the SQL scripts for the database migrations.
- `internal`: Contains the core logic of the application.
  - `delivery`: Handles the HTTP request handling and routing.
  - `domain`: Defines the core entities and interfaces for the application.
  - `repository`: Implements the data access layer.
  - `usecase`: Implements the main business logic of the application.
- `pkg`: Contains utility packages, such as configuration, database connection, middleware, and logging.

## Getting Started

### Prerequisites

- Go version 1.16 or higher
- Docker and Docker Compose

### Installation

1. Clone the repository:

```
git clone https://github.com/lunadiotic/golang-travel-booking.git
```

2. Change to the project directory:

```
cd golang-travel-booking
```

3. Copy the example environment file and update the values:

```
cp .env.example .env
```

### Running the Application

To run the application, use Docker Compose:

```
docker-compose up
```

This will start the API server, PostgreSQL database, RabbitMQ, and Redis services.

### Running Migrations

To run the database migrations, use the provided Makefile commands:

```
make migration_up
```

To run the down migrations:

```
make migration_down
```

## API Documentation

### Authentication

- `POST /api/v1/auth/register`: Register a new user.
- `POST /api/v1/auth/login`: Login and receive a JWT token.

### User Management

- `GET /api/v1/users/profile`: Fetch the logged-in user's profile.
- `PUT /api/v1/users/profile`: Update the logged-in user's profile.

## Testing

The project includes unit tests for the `user_usecase` and the `AuthMiddleware`. You can run the tests using the following command:

```
go test ./...
```

## Deployment

The project is designed to be easily deployable using Docker. You can use the provided `docker-compose.yml` file to deploy the application and its dependencies.

## Built With

- [Go](https://golang.org/) - The programming language used
- [Gin](https://gin-gonic.com/) - The web framework used
- [GORM](https://gorm.io/) - The ORM library used for database interactions
- [Migrate](https://github.com/golang-migrate/migrate) - The database migration tool
- [Testify](https://github.com/stretchr/testify) - The testing framework used

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you find any problems or have suggestions for improvements.

## License

This project is licensed under the [MIT License](LICENSE).
