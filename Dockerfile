# Use the official Go image as the builder stage
FROM golang:1.21 as builder

# Copy go.mod and go.sum to the WORKDIR
COPY go.mod go.sum /go/src/github.com/shubhamm700/Go-Task-API/

# Set the working directory
WORKDIR /go/src/github.com/shubhamm700/Go-Task-API

# Download dependencies
RUN go mod download

# Copy the entire project to the WORKDIR
COPY . /go/src/github.com/shubhamm700/Go-Task-API

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/Go-Task-API github.com/shubhamm700/Go-Task-API

# Use a minimal Alpine image as the final stage
FROM alpine

# Install CA certificates
RUN apk add --no-cache ca-certificates && update-ca-certificates

# Copy the built binary from the builder stage
COPY --from=builder /go/src/github.com/shubhamm700/Go-Task-API/build/Go-Task-API /usr/bin/Go-Task-API

# Expose the necessary port
EXPOSE 8080

# Set the entry point to run the binary
ENTRYPOINT ["/usr/bin/Go-Task-API"]
