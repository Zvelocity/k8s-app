# Use the official Golang image as a base
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod ./

# Download dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the entire Go application
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from a smaller image
FROM alpine:latest  

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

