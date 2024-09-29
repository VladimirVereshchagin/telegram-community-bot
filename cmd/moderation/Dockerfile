# Используем официальный образ Golang в качестве базового
FROM golang:1.20-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Компилируем приложение
RUN go build -o moderation ./cmd/moderation

# Указываем команду запуска
CMD ["./moderation"]