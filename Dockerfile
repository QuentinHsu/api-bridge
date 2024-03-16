# First stage: build the executable
FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

# Second stage: copy the executable and run it
FROM debian:latest
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /app
COPY --from=builder /app .
CMD ["./main"]