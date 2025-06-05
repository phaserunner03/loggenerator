FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o server .

# Use full Debian image for debugging
FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /app/server /server

EXPOSE 8080
CMD ["/server"]
