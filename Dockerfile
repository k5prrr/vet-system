# Dockerfile
FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/main ./cmd/app

# ---
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

# Копируем бинарник и статику
COPY --from=builder /app/bin/main .
COPY --from=builder /app/static ./static
COPY --from=builder /app/.env .env

# Открываем порт
EXPOSE 8080

# Запуск
CMD ["./main"]