# Use an official Go runtime as a parent image
FROM golang:1.18

# Set the working directory in the container
WORKDIR /app

# Copy the Go files
COPY . .

# Build the Go app
RUN go build -o server .

# Expose port 8080 for the application
EXPOSE 8080

# Run the compiled binary
CMD ["./server"]
