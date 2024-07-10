package service

// мок методы сервиса для теста хендлеров

import (
	"github.com/dusk-chancellor/sc_service/internal/models"
)

type MockServiceMethods struct {
}

func (m *MockServiceMethods) SaveOrder(client *models.Client, order *models.HistoryOrder) error {
	return nil
}

func (m *MockServiceMethods) SaveOrderBook(exchange_name, pair string, orderBook []*models.DepthOrder) error {
	return nil
}

func (m *MockServiceMethods) GetOrderBook(exchange_name, pair string) ([]*models.DepthOrder, error) {
	return nil, nil
}

func (m *MockServiceMethods) GetOrderHistory(client *models.Client)  ([]*models.HistoryOrder, error) {
	return nil, nil
}