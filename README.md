# Learning Go REST API From Educative

## Overview
This project is a simple REST API built using Go and the Fiber framework. It demonstrates how to set up a basic web server and define routes.

## Getting Started

### Prerequisites
- Go (version 1.16 or later)
- Fiber framework
- Database setup (e.g., MongoDB, PostgreSQL)

### Installation
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

### Configuration
Before running the application, ensure you have the following environment variable set:
- `DB_NAME`: The name of your database.

You can also set the `PORT` environment variable to specify the port on which the server will run. If not set, it defaults to `8080`.

### Running the Application
To start the server, run:
```
go run main.go
```
## Directory Structure
- `main.go`: The entry point of the application where the Fiber app is initialized and routes are set up.
- `routes/`: Contains route definitions.
- `database/`: Handles database connections and operations.
- `utils/`: Utility functions for configuration and other common tasks.

## License
This project is licensed under the MIT License.