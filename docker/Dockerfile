FROM golang:1.25.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/app

# runtime stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app /app/app
COPY --from=builder /app/migrations /app/migrations
RUN chmod +x /app/app

EXPOSE 8080

CMD ["/app/app"]