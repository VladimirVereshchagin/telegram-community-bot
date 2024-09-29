package automation

import (
	"fmt"
	"log"
	"time"

	botapi "github.com/vladimirvereshchagin/telegram-community-bot/pkg/bot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// AutomationService предоставляет методы для автоматизации задач
type AutomationService struct {
	Bot      botapi.BotAPI // Интерфейс для работы с Telegram Bot API
	stopChan chan struct{} // Канал для остановки планировщика
	ChatID   int64         // Идентификатор чата для отправки сообщений
}

// NewAutomationService создает новый экземпляр AutomationService
func NewAutomationService(bot botapi.BotAPI, chatID int64) *AutomationService {
	return &AutomationService{
		Bot:      bot,
		stopChan: make(chan struct{}),
		ChatID:   chatID,
	}
}

// Start запускает сервис автоматизации
func (a *AutomationService) Start() error {
	// Запуск планировщика задач
	go a.startScheduler()

	// Здесь можно добавить обработку команд и сообщений

	// Бесконечный цикл для поддержания работы сервиса
	select {}
}

// startScheduler запускает периодические задачи
func (a *AutomationService) startScheduler() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			a.sendDailyMessage()
		case <-a.stopChan:
			return
		}
	}
}

// Stop останавливает планировщик
func (a *AutomationService) Stop() {
	close(a.stopChan)
}

// sendDailyMessage отправляет ежедневное сообщение в чат
func (a *AutomationService) sendDailyMessage() {
	msg := tgbotapi.NewMessage(a.ChatID, "Доброе утро!")
	_, err := a.Bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка отправки ежедневного сообщения: %v", err)
	}
}

// HandleFAQ обрабатывает часто задаваемые вопросы
func (a *AutomationService) HandleFAQ(question string) string {
	// Здесь может быть логика поиска ответа на вопрос
	answer := fmt.Sprintf("Вы спросили: %s. К сожалению, у меня пока нет ответа.", question)
	return answer
}
