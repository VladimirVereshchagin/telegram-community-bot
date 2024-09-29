package main

import (
	"log"
	"os"

	"github.com/vladimirvereshchagin/telegram-community-bot/internal/common"
	"github.com/vladimirvereshchagin/telegram-community-bot/internal/user_management"
	botapi "github.com/vladimirvereshchagin/telegram-community-bot/pkg/bot"

	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения из .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки файла .env, будут использоваться переменные окружения")
	}

	// Загружаем конфигурацию
	config, err := common.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Перезаписываем токен из переменных окружения, если он установлен
	if token := os.Getenv("TELEGRAM_TOKEN"); token != "" {
		config.Telegram.Token = token
	}

	// Создаем экземпляр бота
	bot, err := botapi.NewBot(config.Telegram.Token)
	if err != nil {
		log.Fatalf("Ошибка создания бота: %v", err)
	}

	// Создаем сервис управления пользователями
	userService := user_management.NewUserService(bot)

	// Запуск сервиса управления пользователями
	if err := userService.Start(); err != nil {
		log.Fatalf("Ошибка запуска сервиса управления пользователями: %v", err)
	}
}
