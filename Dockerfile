# Build stage
FROM golang:1.24.3-alpine3.21 AS builder

# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Swagger documentation
RUN swag init --parseDependency --parseInternal --output ./docs

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o recipe-api ./cmd/recipe_api

# Final stage
FROM alpine:3.21.3

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/recipe-api .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./recipe-api"]
