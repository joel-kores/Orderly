# Build stage
FROM golang:1.23-alpine AS builder
LABEL authors="Joel Kores"

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o Orderly ./cmd/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary and .env file
COPY --from=builder /app/Orderly .
COPY --from=builder /app/.env .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./Orderly"]