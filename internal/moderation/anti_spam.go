package moderation

import (
	"strings"

	botapi "github.com/vladimirvereshchagin/telegram-community-bot/pkg/bot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// AntiSpam содержит методы для борьбы со спамом
type AntiSpam struct {
	Bot  botapi.BotAPI
	Repo ModerationRepository
}

// NewAntiSpam создает новый экземпляр AntiSpam
func NewAntiSpam(bot botapi.BotAPI, repo ModerationRepository) *AntiSpam {
	return &AntiSpam{
		Bot:  bot,
		Repo: repo,
	}
}

// CheckMessage проверяет сообщение на наличие запрещенных слов
func (a *AntiSpam) CheckMessage(message *tgbotapi.Message) bool {
	blacklistedWords, err := a.Repo.GetBlacklistedWords()
	if err != nil {
		// Обработка ошибки
		return false
	}

	for _, word := range blacklistedWords {
		if strings.Contains(strings.ToLower(message.Text), strings.ToLower(word)) {
			return true
		}
	}
	return false
}

// DeleteMessage удаляет сообщение
func (a *AntiSpam) DeleteMessage(chatID int64, messageID int) error {
	deleteConfig := tgbotapi.DeleteMessageConfig{
		ChatID:    chatID,
		MessageID: messageID,
	}
	_, err := a.Bot.Request(deleteConfig)
	return err
}
