# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Copy go mod files first
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o package-calculator .

# Runtime stage
FROM alpine:latest
WORKDIR /root/

# Copy binary and migrations
COPY --from=builder /app/package-calculator .
COPY --from=builder /app/migrations ./migrations

# Copy configuration files (if any)
COPY ../.env .


EXPOSE 8080
CMD ["./package-calculator"]