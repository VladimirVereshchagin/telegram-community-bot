package analytics

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vladimirvereshchagin/telegram-community-bot/internal/common"
	"github.com/vladimirvereshchagin/telegram-community-bot/internal/moderation"
)

// HTTPClient интерфейс для HTTP клиента (для возможного мокирования)
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// AnalyticsService представляет Google Analytics API
type AnalyticsService struct {
	MeasurementID string     // Идентификатор измерения Google Analytics
	APISecret     string     // Секретный ключ API Google Analytics
	HTTPClient    HTTPClient // HTTP клиент для отправки запросов
}

// Проверяем, что AnalyticsService реализует интерфейс moderation.AnalyticsServiceInterface
var _ moderation.AnalyticsServiceInterface = (*AnalyticsService)(nil)

// NewAnalyticsService создает новый экземпляр AnalyticsService
func NewAnalyticsService(measurementID, apiSecret string) *AnalyticsService {
	return &AnalyticsService{
		MeasurementID: measurementID,
		APISecret:     apiSecret,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second, // Устанавливаем таймаут для HTTP клиента
		},
	}
}

// TrackEvent отправляет событие в Google Analytics
func (a *AnalyticsService) TrackEvent(userID int64, eventName string, params map[string]interface{}) {
	url := fmt.Sprintf(
		"https://www.google-analytics.com/mp/collect?measurement_id=%s&api_secret=%s",
		a.MeasurementID,
		a.APISecret,
	)

	// Генерируем анонимный ClientID на основе хешированного UserID
	clientID := hashUserID(userID)

	// Параметры события
	eventData := map[string]interface{}{
		"client_id": clientID,
		"events": []map[string]interface{}{
			{
				"name":   eventName,
				"params": params,
			},
		},
	}

	// Преобразуем данные в JSON
	jsonData, err := json.Marshal(eventData)
	if err != nil {
		common.HandleError(err, "Ошибка кодирования JSON")
		return
	}

	// Создаем HTTP-запрос
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		common.HandleError(err, "Ошибка создания HTTP-запроса")
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Отправляем запрос
	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		common.HandleError(err, "Ошибка отправки данных в Google Analytics")
		return
	}
	defer resp.Body.Close()

	// Логируем результат отправки
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent {
		log.Println("Событие успешно отправлено в Google Analytics.")
	} else {
		log.Printf("Не удалось отправить событие: статус код %d", resp.StatusCode)
	}
}

// hashUserID хеширует UserID с использованием SHA-256 и возвращает строковое представление
func hashUserID(userID int64) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d", userID)))
	return hex.EncodeToString(h.Sum(nil))
}
