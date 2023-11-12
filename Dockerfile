# Compile stage
FROM golang:1.19 AS build-env

# Install Delve debugger
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Set the working directory inside the container
WORKDIR /dockerdev

# Copy go.mod and go.sum files to download dependencies
COPY ./backend/go.mod ./backend/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY ./backend/ .

# Build the application with debugging flags
RUN go build -gcflags="all=-N -l" -o /server ./cmd/api

# Final stage
FROM debian:buster

# Expose application and Delve ports
EXPOSE 8080 40000

# Copy the Delve debugger and compiled application from the build stage
COPY --from=build-env /go/bin/dlv /
COPY --from=build-env /server /

# Start the application with Delve
CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/server"]
