package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/vladimirvereshchagin/telegram-community-bot/internal/automation"
	"github.com/vladimirvereshchagin/telegram-community-bot/internal/common"
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
