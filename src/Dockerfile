# Base image
FROM golang:1.22

# Add our own layers
WORKDIR /go/src/app

# Install libraries
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download && go mod verify

# Copy the source code of app
COPY ./ ./

# Build the app
RUN go build -o /go/bin/app

# Run the app
CMD ["/go/bin/app"]