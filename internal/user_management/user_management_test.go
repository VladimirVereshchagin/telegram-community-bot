package user_management

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

// MockBot используется для мокирования BotAPI
type MockBot struct {
	Messages []tgbotapi.Chattable
}

// Send мокирует отправку сообщений
func (m *MockBot) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	m.Messages = append(m.Messages, c)
	return tgbotapi.Message{}, nil
}

func (m *MockBot) Request(c tgbotapi.Chattable) (*tgbotapi.APIResponse, error) {
	return &tgbotapi.APIResponse{Ok: true}, nil
}

func (m *MockBot) GetUpdatesChan() tgbotapi.UpdatesChannel {
	return make(chan tgbotapi.Update)
}

func TestUserService_handleNewMember(t *testing.T) {
	bot := &MockBot{}
	service := NewUserService(bot)

	chatID := int64(-1001234567890)
	user := tgbotapi.User{ID: 12345, FirstName: "John"}

	service.handleNewMember(chatID, user)

	// Проверяем, что приветственное сообщение было отправлено
	assert.Len(t, bot.Messages, 1)
	msg, ok := bot.Messages[0].(tgbotapi.MessageConfig)
	assert.True(t, ok)
	assert.Equal(t, chatID, msg.ChatID)
	assert.Equal(t, "Добро пожаловать, John!", msg.Text)
}

func TestUserService_handleLeftMember(t *testing.T) {
	bot := &MockBot{}
	service := NewUserService(bot)

	chatID := int64(-1001234567890)
	user := tgbotapi.User{ID: 12345, FirstName: "John"}

	service.handleLeftMember(chatID, user)

	// Проверяем, что сообщение о выходе участника было отправлено
	assert.Len(t, bot.Messages, 1)
	msg, ok := bot.Messages[0].(tgbotapi.MessageConfig)
	assert.True(t, ok)
	assert.Equal(t, chatID, msg.ChatID)
	assert.Equal(t, "John покинул(а) чат.", msg.Text)
}
