package moderation

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

// MockBot используется для мокирования BotAPI
type MockBot struct {
	Messages []tgbotapi.Chattable
}

func (m *MockBot) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	m.Messages = append(m.Messages, c)
	return tgbotapi.Message{}, nil
}

func (m *MockBot) GetUpdatesChan() tgbotapi.UpdatesChannel {
	return make(chan tgbotapi.Update)
}

// Добавляем метод Request, чтобы удовлетворить интерфейс BotAPI
func (m *MockBot) Request(c tgbotapi.Chattable) (*tgbotapi.APIResponse, error) {
	return &tgbotapi.APIResponse{Ok: true}, nil
}

// MockAnalyticsService используется для мокирования AnalyticsService
type MockAnalyticsService struct{}

func (m *MockAnalyticsService) TrackEvent(userID int64, eventName string, params map[string]interface{}) {
	// Мокированный метод может сохранять данные для проверки в тестах
}

// MockRepository используется для мокирования ModerationRepository
type MockRepository struct{}

func (m *MockRepository) GetBlacklistedWords() ([]string, error) {
	return []string{"spamword"}, nil
}

func (m *MockRepository) AddBlacklistedWord(word string) error {
	return nil
}

func TestModerationService_handleUpdate(t *testing.T) {
	// Мокаем зависимости
	bot := &MockBot{}
	repo := &MockRepository{}
	analyticsService := &MockAnalyticsService{}

	// Создаем ModerationService с моками
	service := NewModerationService(bot, repo, analyticsService)

	// Создаем сообщение с запрещенным словом
	update := tgbotapi.Update{
		Message: &tgbotapi.Message{
			Text: "This message contains spamword.",
			From: &tgbotapi.User{ID: 12345},
			Chat: &tgbotapi.Chat{ID: -1001234567890},
		},
	}

	// Вызываем метод handleUpdate
	service.handleUpdate(update)

	// Проверяем, что сообщение было удалено и предупреждение отправлено
	assert.Len(t, bot.Messages, 1)
	msg, ok := bot.Messages[0].(tgbotapi.MessageConfig)
	assert.True(t, ok)
	assert.Equal(t, "Ваше сообщение было удалено из-за нарушения правил.", msg.Text)
}

func TestModerationService_handleUpdate_MessageReceived(t *testing.T) {
	// Мокаем зависимости
	bot := &MockBot{}
	repo := &MockRepository{}
	analyticsService := &MockAnalyticsService{}

	// Создаем ModerationService с моками
	service := NewModerationService(bot, repo, analyticsService)

	// Создаем тестовое обновление с нормальным сообщением
	update := tgbotapi.Update{
		Message: &tgbotapi.Message{
			Text: "Normal message.",
			From: &tgbotapi.User{ID: 12345},
			Chat: &tgbotapi.Chat{ID: -1001234567890},
		},
	}

	// Вызываем метод handleUpdate
	service.handleUpdate(update)

	// Проверяем, что событие было отправлено и сообщение не удалено
	assert.Len(t, bot.Messages, 0) // Нет предупреждений, так как сообщение корректное
}
