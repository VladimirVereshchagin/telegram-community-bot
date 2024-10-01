# Используем официальный образ Golang в качестве базового
FROM golang:1.22-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum из корневой директории
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY ./cmd/user_management/ ./cmd/user_management/
COPY ./internal/ ./internal/
COPY ./pkg/ ./pkg/

# Компилируем приложение
RUN go build -o user_management ./cmd/user_management

# Указываем команду для запуска
CMD ["./user_management"]