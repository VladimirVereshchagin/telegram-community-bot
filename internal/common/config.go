package common

import (
	"os"

	"gopkg.in/yaml.v2"
)

// DatabaseConfig хранит параметры подключения к базе данных
type DatabaseConfig struct {
	Driver string `yaml:"driver"`
	DSN    string `yaml:"dsn"`
}

// Config структура для хранения конфигурации приложения
type Config struct {
	Telegram struct {
		Token string `yaml:"token"`
	} `yaml:"telegram"`
	Database  DatabaseConfig `yaml:"database"`
	Analytics struct {
		MeasurementID string `yaml:"measurement_id"`
		APISecret     string `yaml:"api_secret"`
	} `yaml:"analytics"`
}

// LoadConfig загружает конфигурацию из YAML файла
func LoadConfig(path string) (*Config, error) {
	// Читаем файл конфигурации
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Расширяем переменные окружения в содержимом файла
	data = []byte(os.ExpandEnv(string(data)))

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
