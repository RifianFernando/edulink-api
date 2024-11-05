# Use the official Go image with a specific version
FROM golang:1.23.1-alpine as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy

# Copy the rest of the application code to the working directory
COPY . .

COPY .env.example .env

# Accept build arguments
ARG APP_KEY
ARG DB_HOST
ARG DB_USERNAME
ARG DB_PASSWORD
ARG ENV_PORT
ARG ENV_ALLOW_ORIGIN

# Set the environment variables
ENV APP_KEY=$APP_KEY
ENV DB_HOST=$DB_HOST
ENV DB_USERNAME=$DB_USERNAME
ENV DB_PASSWORD=$DB_PASSWORD
ENV ENV_PORT=$ENV_PORT
ENV ENV_ALLOW_ORIGIN=$ENV_ALLOW_ORIGIN

# generate the session key
RUN go run database/database.go -key:generate

# Build the Go application
RUN go build -o main ./main.go

# Start a new stage from a lightweight base image
FROM alpine:latest

# Set the working directory in the new image
WORKDIR /app

# Copy the binary from the builder stage to the new image
COPY --from=builder /app/main .

# Make the binary executable (this may not be necessary since it's already done in the builder stage)
RUN chmod +x main

# Expose the port the application runs on
EXPOSE 8000

# Command to run the application
CMD ["./main"]
