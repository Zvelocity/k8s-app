# Build stage
FROM golang:1.24-alpine AS build

WORKDIR /app

# Copy go mod and sum files
COPY src/go.mod src/go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY src/ ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-web-app

# Final stage
FROM alpine:3.19

WORKDIR /

# Copy the binary from the build stage
COPY --from=build /go-web-app /go-web-app

# Expose port
EXPOSE 8080

# Command to run
CMD ["/go-web-app"]
