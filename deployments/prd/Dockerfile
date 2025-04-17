# ---- Build Stage ----
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod and go.sum before other files for better cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary statically
# CGO_ENABLED=0: disable C linking (no glibc/musl issues)
# GOARCH=amd64: target x86_64 (for Intel/AMD)
# GOOS=linux: target Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -o main .



# ---- Final Stage ----
FROM ubuntu:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the .env file to the runtime image
COPY --from=builder /app/.env .

# Ensure the binary is executable
RUN chmod +x ./main

# Command to run
CMD ["./main"]