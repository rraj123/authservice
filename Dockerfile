# Use the official golang image as the base image for building
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files (if they exist)
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
# CGO_ENABLED=0 creates a statically linked binary
# -o specifies the output binary name
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

# Use a minimal alpine image for the final stage
FROM alpine:3.19

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/api .

# Expose the port the application runs on
EXPOSE 4000

# Command to run the application
# The flags can be overridden using docker run command
CMD ["./api", "-port=4000", "-env=development"]