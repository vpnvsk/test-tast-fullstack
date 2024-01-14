
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files into the container
COPY . /app

# Download and install any dependencies
RUN go mod download

# Build the backend application
RUN go build cmd/main.go 

# Expose the port that the application will run on
EXPOSE 8000

# Command to run the backend application
CMD ["./main"]

