package common

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq" // Драйвер для PostgreSQL
)

// NewDatabase создает новое подключение к базе данных
func NewDatabase(config DatabaseConfig) (*sql.DB, error) {
	// Расширяем переменные окружения в DSN
	dsn := os.ExpandEnv(config.DSN)

	db, err := sql.Open(config.Driver, dsn)
	if err != nil {
		return nil, err
	}
	// Проверка подключения
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
