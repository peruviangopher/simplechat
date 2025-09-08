# Stage 1: Build
FROM golang:1.23 AS builder

# Set working directory
WORKDIR /app

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the Go app (static binary, no CGO)
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./main.go

# Stage 2: Run
FROM alpine:latest

# Create a non-root user (optional but recommended)
RUN adduser -D appuser

WORKDIR /app

COPY ./templates app/templates

# Copy binary from builder
COPY --from=builder /app/server .

COPY --from=builder /app/templates ./templates

# Expose the Gin port
EXPOSE 8080

# Run as non-root user
USER appuser

CMD ["./server", "--rooms", "3"]