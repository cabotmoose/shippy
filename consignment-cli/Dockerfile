# Use official golang image
# `as builder` gives this container a name that can be referenced later
FROM golang:1.10.0 as builder

# Set our workdir to current service
WORKDIR /go/src/github.com/cabotmoose/shippy/consignment-cli

# Copy current code into workdir
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# Second FROM statement tells Docker to start new build process
FROM alpine:latest

# Add security package
RUN apk --no-cache add ca-certificates

# Create directory for app
RUN mkdir /app
WORKDIR /app

# Copy JSON file
COPY consignment.json /app/consignment.json

# Instead of copying binary from our host machine, we pull binary from container named `builder`
COPY --from=builder /go/src/github.com/cabotmoose/shippy/consignment-cli/consignment-cli .

# Run the binary
CMD ["./consignment-cli"]