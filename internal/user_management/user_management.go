package user_management

import (
	"fmt"
	"log"

	botapi "github.com/vladimirvereshchagin/telegram-community-bot/pkg/bot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// UserService предоставляет методы для управления пользователями
type UserService struct {
	Bot botapi.BotAPI
}

// NewUserService создает новый экземпляр UserService
func NewUserService(bot botapi.BotAPI) *UserService {
	return &UserService{
		Bot: bot,
	}
}

// Start запускает сервис управления пользователями
func (u *UserService) Start() error {
	// Устанавливаем команды бота
	u.setupBotCommands()

	updates := u.Bot.GetUpdatesChan()
	for update := range updates {
		u.handleUpdate(update)
	}
	return nil
}

// setupBotCommands устанавливает команды бота
func (u *UserService) setupBotCommands() {
	commands := []tgbotapi.BotCommand{
		{Command: "start", Description: "Начать работу с ботом"},
		{Command: "help", Description: "Справка по использованию бота"},
	}

	cfg := tgbotapi.NewSetMyCommands(commands...)
	_, err := u.Bot.Request(cfg)
	if err != nil {
		log.Printf("Ошибка установки команд бота: %v", err)
	}
}

// handleUpdate обрабатывает обновления
func (u *UserService) handleUpdate(update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		u.handleCallbackQuery(update.CallbackQuery)
		return
	}

	if update.Message == nil {
		return
	}

	log.Printf("Получено обновление: %+v", update)

	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			u.handleStartCommand(update.Message)
		case "help":
			u.handleHelpCommand(update.Message)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Используйте /help для списка доступных команд.")
			u.Bot.Send(msg)
		}
		return
	}

	// Обработка новых участников
	if len(update.Message.NewChatMembers) > 0 {
		for _, newUser := range update.Message.NewChatMembers {
			u.handleNewMember(update.Message.Chat.ID, newUser)
		}
	}

	// Обработка уходящих участников
	if update.Message.LeftChatMember != nil {
		u.handleLeftMember(update.Message.Chat.ID, *update.Message.LeftChatMember)
	}
}

// handleStartCommand обрабатывает команду /start
func (u *UserService) handleStartCommand(message *tgbotapi.Message) {
	welcomeMsg := fmt.Sprintf("Здравствуйте, %s! Добро пожаловать в наше сообщество.", message.From.FirstName)
	msg := tgbotapi.NewMessage(message.Chat.ID, welcomeMsg)
	_, err := u.Bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка отправки сообщения: %v", err)
	}
}

// handleHelpCommand обрабатывает команду /help
func (u *UserService) handleHelpCommand(message *tgbotapi.Message) {
	// Создаем инлайн-клавиатуру
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("FAQ", "faq"),
			tgbotapi.NewInlineKeyboardButtonData("Правила сообщества", "rules"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Связаться с поддержкой", "support"),
			tgbotapi.NewInlineKeyboardButtonData("Настройки", "settings"),
		),
	)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите опцию:")
	msg.ReplyMarkup = keyboard

	_, err := u.Bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка отправки сообщения с инлайн-клавиатурой: %v", err)
	}
}

// handleCallbackQuery обрабатывает нажатия на инлайн-кнопки
func (u *UserService) handleCallbackQuery(callback *tgbotapi.CallbackQuery) {
	var responseMsg string

	switch callback.Data {
	case "faq":
		responseMsg = "Часто задаваемые вопросы:\n1. Как пользоваться ботом?\n2. Как связаться с поддержкой?"
	case "rules":
		responseMsg = "Правила сообщества:\n1. Соблюдайте вежливость.\n2. Запрещен спам и реклама."
	case "support":
		responseMsg = "Для связи с поддержкой напишите на email: support@example.com"
	case "settings":
		responseMsg = "Настройки пока недоступны. Функция в разработке."
	default:
		responseMsg = "Извините, я не понимаю эту команду."
	}

	// Отправляем ответ пользователю
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, responseMsg)
	_, err := u.Bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка отправки ответа на CallbackQuery: %v", err)
	}

	// Уведомляем Telegram, что CallbackQuery обработан
	callbackConfig := tgbotapi.NewCallback(callback.ID, "")
	if _, err := u.Bot.Request(callbackConfig); err != nil {
		log.Printf("Ошибка отправки Callback: %v", err)
	}
}

// handleNewMember обрабатывает новых участников
func (u *UserService) handleNewMember(chatID int64, user tgbotapi.User) {
	welcomeMsg := fmt.Sprintf("Добро пожаловать, %s!", user.FirstName)
	msg := tgbotapi.NewMessage(chatID, welcomeMsg)
	_, err := u.Bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка отправки приветственного сообщения: %v", err)
	} else {
		log.Println("Приветственное сообщение отправлено")
	}
}

// handleLeftMember обрабатывает уходящих участников
func (u *UserService) handleLeftMember(chatID int64, user tgbotapi.User) {
	farewellMsg := fmt.Sprintf("%s покинул(а) чат.", user.FirstName)
	msg := tgbotapi.NewMessage(chatID, farewellMsg)
	_, err := u.Bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка отправки сообщения о выходе участника: %v", err)
	}
}
