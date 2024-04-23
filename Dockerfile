# Use the official Go image to create a build artifact.
FROM golang:1.21.2-alpine as builder

# Set the working directory inside the container.
WORKDIR /app/backend

# Copy go mod and sum files.
COPY ./backend/go.* ./

# Download dependencies.
RUN go mod download

# Copy the backend source files and build it.
COPY ./backend/ ./

RUN go build -o main .

# Final stage: Create the runtime container.
FROM alpine:latest as runtime

# Install ca-certificates in case you need HTTPS.
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the compiled backend application from the builder image.
COPY --from=builder /app/backend/main /app

# Copy the frontend files from the local frontend directory.
COPY ./frontend /app/frontend

# Expose port 8080.
EXPOSE 8080

# Command to run the executable.
CMD ["/app/main"]