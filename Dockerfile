# Use the official Golang image as the base image
FROM golang:1.25-alpine as builder
WORKDIR /app
COPY . .
RUN go build ./cmd/redirect

# Use slim alpine image for production
FROM alpine:3.23 as production
WORKDIR /app
COPY --from=builder /app/redirect .
EXPOSE 3000

# Run the Go program when the container starts
CMD ["./redirect"]
