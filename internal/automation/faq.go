package automation

// Здесь могут быть функции и данные, связанные с FAQ

// FAQItem представляет вопрос и ответ
type FAQItem struct {
	Question string
	Answer   string
}

// FAQList содержит список часто задаваемых вопросов
var FAQList = []FAQItem{
	{"Как воспользоваться ботом?", "Вы можете использовать команды /start, /help и другие."},
	{"Как связаться с поддержкой?", "Напишите на support@example.com."},
}

// GetFAQAnswer ищет ответ на заданный вопрос
func GetFAQAnswer(question string) string {
	for _, item := range FAQList {
		if item.Question == question {
			return item.Answer
		}
	}
	return "К сожалению, я не нашел ответ на ваш вопрос."
}
