package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vladimirvereshchagin/telegram-community-bot/internal/analytics"
	"github.com/vladimirvereshchagin/telegram-community-bot/internal/common"
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

	// Перезаписываем параметры аналитики из переменных окружения
	if measurementID := os.Getenv("MEASUREMENT_ID"); measurementID != "" {
		config.Analytics.MeasurementID = measurementID
	}
	if apiSecret := os.Getenv("API_SECRET"); apiSecret != "" {
		config.Analytics.APISecret = apiSecret
	}

	// Создаем экземпляр AnalyticsService
	analyticsService := analytics.NewAnalyticsService(
		config.Analytics.MeasurementID,
		config.Analytics.APISecret,
	)

	// Логируем создание сервиса для проверки использования переменной
	log.Println("Сервис аналитики успешно инициализирован:", analyticsService)
}
