package common

import (
	"database/sql"

	_ "github.com/lib/pq" // Драйвер для PostgreSQL
)

// NewDatabase создает новое подключение к базе данных
func NewDatabase(config DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open(config.Driver, config.DSN)
	if err != nil {
		return nil, err
	}
	// Проверка подключения
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
