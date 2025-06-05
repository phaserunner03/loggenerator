# Stage 1: Build the Go binary with Go 1.22+
FROM golang:1.21 as builder

WORKDIR /app
COPY . .

# Make sure go.mod is compatible with Go 1.22 (not 1.24)
RUN go mod tidy
RUN go build -o server .

# Stage 2: Minimal runtime image (with CA certs and no shell)
FROM gcr.io/distroless/base-debian11

COPY --from=builder /app/server /server

# Cloud Run expects it to listen on $PORT
CMD ["/server"]
