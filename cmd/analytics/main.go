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
