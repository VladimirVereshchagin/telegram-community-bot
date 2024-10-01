# Используем официальный образ Golang с поддержкой Alpine
FROM golang:1.22-alpine

# Создаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем все необходимые исходные файлы (включая internal и pkg)
COPY ./cmd/moderation/ ./cmd/moderation/
COPY ./internal/ ./internal/
COPY ./pkg/ ./pkg/

# Компилируем приложение
RUN go build -o moderation ./cmd/moderation

# Указываем команду запуска
CMD ["./moderation"]
