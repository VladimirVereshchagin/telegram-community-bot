package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// BotAPI интерфейс для работы с Telegram Bot API
type BotAPI interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
	GetUpdatesChan() tgbotapi.UpdatesChannel
	Request(c tgbotapi.Chattable) (*tgbotapi.APIResponse, error)
}

// Bot реализация интерфейса BotAPI
type Bot struct {
	API *tgbotapi.BotAPI
}

// NewBot создает новый экземпляр Bot
func NewBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Bot{API: api}, nil
}

// Send отправляет сообщение через Bot API
func (b *Bot) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	return b.API.Send(c)
}

// GetUpdatesChan возвращает канал обновлений
func (b *Bot) GetUpdatesChan() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	u.AllowedUpdates = []string{"message", "my_chat_member", "chat_member"}
	updates := b.API.GetUpdatesChan(u)
	return updates
}

// Request отправляет произвольный запрос к Telegram API
func (b *Bot) Request(c tgbotapi.Chattable) (*tgbotapi.APIResponse, error) {
	return b.API.Request(c)
}
