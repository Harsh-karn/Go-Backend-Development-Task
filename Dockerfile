# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server/main.go

# Run stage
FROM alpine:3.18

WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/server .

# Expose port
EXPOSE 3000

# Command to run
CMD ["/app/server"]
