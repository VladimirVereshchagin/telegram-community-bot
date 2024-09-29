package main

import (
	"log"
	"os"
	"strconv"

	"github.com/vladimirvereshchagin/telegram-community-bot/internal/automation"
	"github.com/vladimirvereshchagin/telegram-community-bot/internal/common"
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

	// Получаем ChatID из переменных окружения
	chatIDStr := os.Getenv("CHAT_ID")
	if chatIDStr == "" {
		log.Fatal("Переменная окружения CHAT_ID не установлена")
	}

	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatalf("Ошибка преобразования CHAT_ID: %v", err)
	}

	// Создаем сервис автоматизации
	automationService := automation.NewAutomationService(bot, chatID)

	// Запуск сервиса автоматизации
	if err := automationService.Start(); err != nil {
		log.Fatalf("Ошибка запуска сервиса автоматизации: %v", err)
	}
}
