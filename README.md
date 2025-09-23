# Car Inventory Management System

This project is a simple **Car Inventory Management System** built with Go (Golang). It demonstrates a modular structure for handling car data, HTTP requests, middleware, and configuration.

## Project Structure

```
carInventory/
│
├── main.go
├── config/
│   └── config.go          # Database configuration
├── models/
│   └── car.go             # Car model structure
├── handlers/
│   └── car_handlers.go    # Handles HTTP requests and responses for cars
├── middleware/
│   ├── logging.go         # Request logging middleware
│   └── security.go        # Security checks before requests reach the server
└── utils/
    └── response.go        # JSON encoding and decoding utilities
```

## Features

- **RESTful API** for managing car inventory
- **Modular codebase** with clear separation of concerns
- **Middleware** for logging and security
- **Configurable database connection**
- **Models** for easy data handling
- **Utility functions** for JSON processing

## Getting Started

### Prerequisites

- Go 1.18+
- (Optional) A running database instance (if using persistent storage)

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/Aashritha123-lab/GO_BASICS.git
    cd GO_BASICS/car_Inventory
    ```

2. Install dependencies (if any):
    ```bash
    go mod tidy
    ```

3. Update the database configuration in `config/config.go` as needed.

### Running the Project

```bash
go run main.go
```

The server will start on the configured port (see your `main.go` or `config.go`).

### API Endpoints

Typical endpoints may include (check your actual implementation for details):

- `GET /cars` — List all cars
- `GET /cars/{id}` — Get details about a specific car
- `POST /cars` — Add a new car
- `PUT /cars/{id}` — Update a car's information
- `DELETE /cars/{id}` — Remove a car from inventory

## Folder Details

- **`main.go`**: Entry point of the application.
- **`config/config.go`**: Database and environment configuration.
- **`models/car.go`**: Car struct and data-related logic.
- **`handlers/car_handlers.go`**: HTTP handlers for car routes.
- **`middleware/`**: Logging and security checks.
- **`utils/response.go`**: Helper functions for response formatting.


---

*Created by [Aashritha123-lab](https://github.com/Aashritha123-lab)*
