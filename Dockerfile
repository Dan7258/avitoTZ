# Dockerfile
FROM golang:1.25.1-alpine

# Иногда нужен git/ca-certificates для приватных модулей или TLS
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# Кэшируем зависимости
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Копируем весь код
COPY . .

# Открываем порт, как требует ТЗ
EXPOSE 8080

# Запускаем прямо из исходников (без предварительной сборки бинарника)
# Это самый простой и надёжный вариант для тестового задания
CMD ["go", "run", "./cmd/avito"]