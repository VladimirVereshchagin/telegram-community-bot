package moderation

import (
	"database/sql"
)

// ModerationRepository определяет методы для доступа к данным модерации
type ModerationRepository interface {
	GetBlacklistedWords() ([]string, error)
	AddBlacklistedWord(word string) error
}

// SQLModerationRepository реализация ModerationRepository с использованием SQL базы данных
type SQLModerationRepository struct {
	db *sql.DB
}

// NewSQLModerationRepository создает новый экземпляр SQLModerationRepository
func NewSQLModerationRepository(db *sql.DB) *SQLModerationRepository {
	return &SQLModerationRepository{
		db: db,
	}
}

// GetBlacklistedWords возвращает список запрещенных слов
func (r *SQLModerationRepository) GetBlacklistedWords() ([]string, error) {
	rows, err := r.db.Query("SELECT word FROM blacklisted_words")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var words []string
	for rows.Next() {
		var word string
		if err := rows.Scan(&word); err != nil {
			return nil, err
		}
		words = append(words, word)
	}
	return words, nil
}

// AddBlacklistedWord добавляет запрещенное слово
func (r *SQLModerationRepository) AddBlacklistedWord(word string) error {
	_, err := r.db.Exec("INSERT INTO blacklisted_words (word) VALUES ($1)", word)
	return err
}
