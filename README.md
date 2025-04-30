```markdown
# Package Calculator

## About the Project

Package Calculator is a Go-based application designed to manage item pack sizes and calculate optimal orders based on 
user input. It provides a RESTful API for backend operations and a user-friendly web interface for 
interacting with the system.

-  **API**: The backend API is available at [https://api.calculator.mikulic.dev](https://api.calculator.mikulic.dev).
-  **UI**: The frontend UI is accessible at [https://calculator.mikulic.dev](https://calculator.mikulic.dev).

## Features

-  Manage pack sizes (add, delete, and view).
-  Calculate optimal orders based on item quantities.
-  Simple and responsive web interface.

## Prerequisites

- Docker
- Docker Compose
- `make` utility

## Getting Started

Follow these steps to set up and run the project locally:

### 1. Clone the Repository

```bash
git clone https://github.com/croatiangrn/package-calculator.git
cd package-calculator
```

### 2. Configure Environment Variables

Ensure the `.env` file is properly configured. The default `.env` file includes:

```dotenv
# Database
DB_USER=myuser
DB_PASSWORD=mypassword
DB_NAME=mydatabase
DB_ROOT_PASSWORD=myrootpassword
DB_HOST=mysql
DB_PORT=3306
APP_PORT=8084
```

### 3. Build and Start the Application

Use the `make` utility to build and start the application:

```bash
make build
make start
```

This command will:
- Build the Go application.
- Start the MySQL database and the application using Docker Compose.

### 4. Stop the Application

To stop the application, run:

```bash
make stop
```

### 5. Access the Application demo online

- **API**: [https://api.calculator.mikulic.dev](https://api.calculator.mikulic.dev)
- **UI**: [https://calculator.mikulic.dev](https://calculator.mikulic.dev)

## Makefile Commands

- `make build`: Build the application.
- `make start`: Start the application.
- `make rebuild`: Rebuild the application and start the containers.
- `make restart`: Restart the application.
- `make stop`: Stop the application.

```