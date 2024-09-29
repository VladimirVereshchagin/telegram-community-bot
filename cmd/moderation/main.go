package main

import (
	"log"
	"os"

	"github.com/vladimirvereshchagin/telegram-community-bot/internal/analytics"
	"github.com/vladimirvereshchagin/telegram-community-bot/internal/common"
	"github.com/vladimirvereshchagin/telegram-community-bot/internal/moderation"
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

	// Создаем подключение к базе данных
	db, err := common.NewDatabase(config.Database)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Создаем репозиторий модерации
	repo := moderation.NewSQLModerationRepository(db)

	// Создаем экземпляр AnalyticsService (используем интерфейс)
	analyticsService := analytics.NewAnalyticsService(
		config.Analytics.MeasurementID,
		config.Analytics.APISecret,
	)

	// Создаем сервис модерации
	moderationService := moderation.NewModerationService(bot, repo, analyticsService)

	// Запуск сервиса модерации
	if err := moderationService.Start(); err != nil {
		log.Fatalf("Ошибка запуска сервиса модерации: %v", err)
	}
}
