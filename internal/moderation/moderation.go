package moderation

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	botapi "github.com/vladimirvereshchagin/telegram-community-bot/pkg/bot"
)

// Определяем интерфейс для AnalyticsService
type AnalyticsServiceInterface interface {
	TrackEvent(userID int64, eventName string, params map[string]interface{})
}

// ModerationService предоставляет методы для модерации контента
type ModerationService struct {
	Bot              botapi.BotAPI
	Repo             ModerationRepository
	AnalyticsService AnalyticsServiceInterface
}

// NewModerationService создает новый экземпляр ModerationService
func NewModerationService(bot botapi.BotAPI, repo ModerationRepository, analyticsService AnalyticsServiceInterface) *ModerationService {
	return &ModerationService{
		Bot:              bot,
		Repo:             repo,
		AnalyticsService: analyticsService,
	}
}

// Start запускает сервис модерации
func (m *ModerationService) Start() error {
	updates := m.Bot.GetUpdatesChan()
	for update := range updates {
		m.handleUpdate(update)
	}
	return nil
}

// handleUpdate обрабатывает обновления Telegram
func (m *ModerationService) handleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	// Проверяем сообщение на спам
	antiSpam := NewAntiSpam(m.Bot, m.Repo)
	if antiSpam.CheckMessage(update.Message) {
		// Удаляем сообщение
		err := antiSpam.DeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
		if err != nil {
			log.Printf("Ошибка удаления сообщения: %v", err)
		}

		// Отправляем предупреждение пользователю
		warningMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваше сообщение было удалено из-за нарушения правил.")
		_, err = m.Bot.Send(warningMsg)
		if err != nil {
			log.Printf("Ошибка отправки предупреждения: %v", err)
		}

		// Отправляем событие в аналитику
		m.AnalyticsService.TrackEvent(update.Message.From.ID, "message_deleted", map[string]interface{}{
			"reason": "blacklisted_word",
		})

		return
	}

	// Отправляем событие в аналитику о полученном сообщении
	m.AnalyticsService.TrackEvent(update.Message.From.ID, "message_received", map[string]interface{}{
		"message_length": len(update.Message.Text),
	})
}
