package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vladimirvereshchagin/telegram-community-bot/internal/common"
	"github.com/vladimirvereshchagin/telegram-community-bot/internal/user_management"
	botapi "github.com/vladimirvereshchagin/telegram-community-bot/pkg/bot"
)

func main() {
	// Загружаем переменные окружения из .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Файл .env не найден, используются переменные окружения из системы")
	}

	// Получаем путь к файлу конфигурации из переменной окружения
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		// Если переменная не установлена, используем путь по умолчанию
		configPath = "configs/config.yaml"
	}

	// Загружаем конфигурацию
	config, err := common.LoadConfig(configPath)
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
