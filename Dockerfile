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

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Copy the .env file
COPY .env ./

# Copy the config directory
COPY config ./config/

# Copy the database directory
COPY database ./database/

# Copy the helper directory
COPY helper ./helper/

# Copy the main.go file
COPY main.go ./

# Copy the request directory
COPY request ./request/

# Copy the connections directory
COPY connections ./connections/

# Copy the lib directory
COPY lib ./lib/

# Copy the lib directory
COPY res ./res/

# Copy the middleware directory
COPY middleware ./middleware/

# Copy the routes directory
COPY routes ./routes/

# Copy the controllers directory
COPY controllers ./controllers/

# Copy the public directory
COPY public ./public/

# Build the Go application & make it executable by changing the permissions
RUN go build -o main ./main.go && chmod +x main

# Expose the port the application runs on
EXPOSE 443

# Command to run the application
CMD ["./main"]
