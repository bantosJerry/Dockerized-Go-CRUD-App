# Use official Golang image as the base
FROM golang:1.22-alpine AS base

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency management
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Install dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o app .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./app"]

# Testing stage
FROM base AS test

# Set environment variables for the test stage
ENV DB_HOST=db
ENV DB_USER=user
ENV DB_PASSWORD=password
ENV DB_NAME=go_crud_db
ENV DB_PORT=5432

# Run tests
CMD ["go", "test", "./...", "-v"]
