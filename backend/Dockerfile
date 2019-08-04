FROM golang:alpine as builder
WORKDIR /app

# Install SSL and git (needed for go mod)
RUN apk update && apk add --no-cache build-base git ca-certificates && update-ca-certificates

# Download and cache deps
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod verify

# Copy source code
COPY . .
# Build the binary
# Disabling CGO: see https://github.com/golang/go/issues/26988
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o goapp

#
# Small final image
#
# FROM scratch
FROM alpine
WORKDIR /app
# Copy SSL certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Copy go binary
COPY --from=builder /app .

EXPOSE 8080
ENV GIN_MODE release
ENTRYPOINT ["/app/goapp"]