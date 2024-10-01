# Используем официальный образ Golang в качестве базового
FROM golang:1.22-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum из корневой директории проекта
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код сервиса automation
COPY cmd/automation/ ./

# Компилируем приложение
RUN go build -o automation .

# Указываем команду запуска
CMD ["./automation"]