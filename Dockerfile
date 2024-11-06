# Use the official Go image with a specific version
FROM golang:1.23.1-alpine

# Create a group and user
RUN addgroup -S edulinkgroup && adduser -S edulink -G edulinkgroup

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Change ownership of the files to the non-root user (edulink)
RUN chown -R edulink:edulinkgroup /app

# Tell Docker that all future commands should run as the edulink user
USER edulink

# Install dependencies
RUN go mod tidy

# Copy the rest of the application code to the working directory
COPY . .

# Build the Go application
RUN go build -o main ./main.go

# Make the binary executable (this may not be necessary since it's already done in the builder stage)
RUN chmod +x main

# Expose the port the application runs on
EXPOSE 8000

# Command to run the application
CMD ["./main"]
