# # Use an official Golang runtime as a parent image
# FROM golang:1.17

# # Set the working directory to /app
# WORKDIR /app

# # Copy the current directory contents into the container at /app
# COPY . /app

# # Install any needed packages specified in go.mod
# RUN go get -d -v

# # Build the app
# RUN go build -o main .

# # Expose port 8080 for the app
# EXPOSE 8080

# # Run the app when the container launches
# CMD ["/app/main"]


# Use the official Golang Alpine image version 1.18 as the base image
FROM golang:alpine AS build

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any needed packages specified in go.mod
RUN go get -d -v

# Build the app
RUN go build -o main .

# Use an Alpine image as the base image for the final stage
FROM alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=build /app/main /app/

# Expose port 8080 for the app
EXPOSE 8080

# Run the app when the container launches
CMD ["/app/main"]