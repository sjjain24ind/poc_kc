# Build stage
FROM golang:1.20 AS builder

# Install musl-cross for cross-compilation
RUN apt-get update && apt-get install -y musl-tools

# Set the working directory
WORKDIR /app

# Copy the Go module files and login.html
COPY go.mod go.sum login.html ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application for Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o admin-app .

# Final stage
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /app/admin-app /admin-app

# Copy the login.html file
COPY --from=builder /app/login.html login.html

# Expose the application port
EXPOSE 9081

# Run the application
CMD ["/admin-app"]
