
# Dockerized Go CRUD Application
A simple CRUD application built with Go, using Docker for containerization. It includes user registration, login, and management features, backed by a PostgreSQL database.

A PostgreSQL database is used to persist user data, while the Go application provides API endpoints for CRUD operations. The application is containerized using Docker, ensuring a consistent development and deployment environment.

## Prerequisites
- Docker  
- Docker Compose  

## Setup
Clone the repository
```bash
   git clone <repo_url>
   cd Dockerized-Go-CRUD-App
```

### Starting the Application and Running Tests
To build and start the application along with the PostgreSQL database, run

```bash
docker-compose up --build
```

This will
- Start the PostgreSQL database on port `5433`.  
- Build and run the Go application on port `8080`.  
- Run tests after the application has started

Visit `http://localhost:8080` to access the application.


## Project Structure
- `Dockerfile`: Defines the container setup for the Go application and testing environments.  
- `docker-compose.yml`: Configuration for running the app, database, and tests as services.  
- `main.go`: The entry point of the Go application.  
- `go.mod` and `go.sum`: Dependency management files.  
- `tests/`: Contains unit and integration tests for the application.

## Notes
- Ensure that Docker and Docker Compose are installed on your system before running the application.  
- Modify the `docker-compose.yml` or `.env` file if you need to change the database credentials or ports.  
- The database service runs on port `5433` (mapped from `5432` inside the container).  
- The Go application runs on port `8080`.  
