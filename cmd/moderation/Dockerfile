# Используем официальный образ Golang в качестве базового
FROM golang:1.22-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum из корневой директории проекта
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код сервиса moderation
COPY cmd/moderation/ ./

# Компилируем приложение
RUN go build -o moderation .

# Указываем команду запуска
CMD ["./moderation"]