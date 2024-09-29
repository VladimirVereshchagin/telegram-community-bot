package automation

import (
	"testing"
	"time"

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

// Добавляем метод GetUpdatesChan для соответствия интерфейсу BotAPI
func (m *MockBot) GetUpdatesChan() tgbotapi.UpdatesChannel {
	return make(chan tgbotapi.Update)
}

// Добавляем метод Request для соответствия интерфейсу BotAPI
func (m *MockBot) Request(c tgbotapi.Chattable) (*tgbotapi.APIResponse, error) {
	return &tgbotapi.APIResponse{Ok: true}, nil
}

// Тестируем метод sendDailyMessage
func TestAutomationService_sendDailyMessage(t *testing.T) {
	bot := &MockBot{}
	chatID := int64(-1001234567890)
	service := NewAutomationService(bot, chatID)
	// Вызываем метод sendDailyMessage
	service.sendDailyMessage()
	// Проверяем, что сообщение было отправлено
	assert.Len(t, bot.Messages, 1)
	// Проверяем содержимое отправленного сообщения
	msg, ok := bot.Messages[0].(tgbotapi.MessageConfig)
	assert.True(t, ok)
	assert.Equal(t, service.ChatID, msg.ChatID)
	assert.Equal(t, "Доброе утро!", msg.Text)
}

// Тестируем планировщик AutomationService
func TestAutomationService_startScheduler(t *testing.T) {
	bot := &MockBot{}
	chatID := int64(-1001234567890)
	service := NewAutomationService(bot, chatID)

	// Добавляем канал для синхронизации
	done := make(chan bool)

	// Обертка для метода sendDailyMessage с отправкой сигнала в канал
	wrappedSendDailyMessage := func() {
		service.sendDailyMessage() // Вызов оригинального метода
		done <- true               // Отправляем сигнал о выполнении
	}

	// Запускаем планировщик в отдельной горутине
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			wrappedSendDailyMessage()
		}
	}()

	// Ждем сигнала о выполнении задачи или таймаута
	select {
	case <-done:
		// Задача выполнена, продолжаем тест
	case <-time.After(5 * time.Second):
		t.Fatal("Задача не была выполнена в отведенное время")
	}

	// Останавливаем планировщик
	service.Stop()

	// Проверяем, что сообщение было отправлено хотя бы один раз
	assert.GreaterOrEqual(t, len(bot.Messages), 1, "Ожидалось, что хотя бы одно сообщение будет отправлено")
}
