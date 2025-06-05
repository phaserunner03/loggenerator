# Stage 1: Build the Go app
FROM golang:1.21 as builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o server .

# Stage 2: Minimal runtime image
FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/server /server

CMD ["/server"]
