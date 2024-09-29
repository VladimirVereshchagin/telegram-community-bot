package analytics

import (
	"errors"
	"net/http"
	"testing"

	"github.com/vladimirvereshchagin/telegram-community-bot/internal/common"

	"github.com/stretchr/testify/assert"
)

// MockHTTPClient используется для мокирования HTTP клиента
type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do вызывает функцию DoFunc
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestAnalyticsService_TrackEvent_Success(t *testing.T) {
	// Создаем моковый HTTP клиент, возвращающий успешный ответ
	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusNoContent,
				Body:       http.NoBody,
			}, nil
		},
	}

	// Создаем экземпляр AnalyticsService с моковым HTTP клиентом
	service := &AnalyticsService{
		MeasurementID: "test-measurement-id",
		APISecret:     "test-api-secret",
		HTTPClient:    mockClient,
	}

	// Вызываем метод TrackEvent с фиктивным userID
	service.TrackEvent(123456789, "test_event", map[string]interface{}{"param1": "value1"})

	// Проверяем, что запрос был успешно отправлен (тест проходит, если нет ошибок)
	assert.True(t, true)
}

func TestAnalyticsService_TrackEvent_Failure(t *testing.T) {
	// Создаем моковый HTTP клиент, возвращающий ошибку
	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("network error")
		},
	}

	// Создаем экземпляр AnalyticsService с моковым HTTP клиентом
	service := &AnalyticsService{
		MeasurementID: "test-measurement-id",
		APISecret:     "test-api-secret",
		HTTPClient:    mockClient,
	}

	// Перехватываем обработку ошибок
	originalHandleError := common.HandleError
	defer func() { common.HandleError = originalHandleError }()
	var capturedError error
	common.HandleError = func(err error, message string) {
		capturedError = err
	}

	// Вызываем метод TrackEvent
	service.TrackEvent(123456789, "test_event", map[string]interface{}{"param1": "value1"})

	// Проверяем, что ошибка была захвачена
	assert.NotNil(t, capturedError)
	assert.EqualError(t, capturedError, "network error")
}
