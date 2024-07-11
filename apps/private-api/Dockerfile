# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source file
COPY src/server.go .

# Build the application
RUN GOOS=linux GOARCH=amd64 go build -o server server.go



# Stage 2: Create the final image with the built application
FROM alpine:latest

# Install tzdata for time zone data
RUN apk add --no-cache tzdata

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Ensure the binary has execution permissions
RUN chmod +x ./server

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the application
CMD ["./server"]