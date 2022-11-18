# Build stage 1
FROM golang:1.18-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Build stage 2
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 3000
CMD ["/app/main"]