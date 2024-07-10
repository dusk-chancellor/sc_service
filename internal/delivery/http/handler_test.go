package http

// тесты для эндпоинтов, структура:
// 1. объявление нужных переменных как данные для реквеста
// 2. структура тестов (данные и ожидаемые респонсы)
// 3. запуск цикла тестов
// все тесты выполнены на проверку статуса 400

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dusk-chancellor/sc_service/internal/models"
	"github.com/dusk-chancellor/sc_service/internal/service"
)

// тест для эндпоинта GET /orderbook
func TestGetOrderBookHandler(t *testing.T) {
	var (
		exchangeName = "test1"
		pair         = "test2"
	)

	tests := []struct {
		name         string
		exchangeName string
		pair         string
		wantCode     int
	}{
		{
			name:         "exchange_name and pair are required",
			exchangeName: "",
			pair:         "",
			wantCode:     http.StatusBadRequest,
		},
		{
			name:         "exchange_name required",
			exchangeName: exchangeName,
			pair:         "",
			wantCode:     http.StatusBadRequest,
		},
		{
			name:         "pair required",
			exchangeName: "",
			pair:         pair,
			wantCode:     http.StatusBadRequest,
		},
		{
			name:         "exchange_name and pair",
			exchangeName: exchangeName,
			pair:         pair,
			wantCode:     http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandlers(&service.MockServiceMethods{})
			req := httptest.NewRequest("GET", "/orderbook?exchange_name="+tt.exchangeName+"&pair="+tt.pair, nil)
			w := httptest.NewRecorder()
			h.GetOrderBookHandler()(w, req)
			if w.Code != tt.wantCode {
				t.Errorf("GetOrderBookHandler() error = %v, wantErr %v", w.Code, tt.wantCode)
			}
		})
	}
}

// тест для эндпоинта POST /orderbook
func TestSaveOrderBookHandler(t *testing.T) {
	var (
		exchangeName = "test1"
		pair         = "test2"
	)

	orderBook := []*models.DepthOrder{
		{
			Price:  1,
			BaseQty: 1,
		},
		{
			Price:  2,
			BaseQty: 2,
		},
	}
	jsonData, err := json.Marshal(orderBook)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name         string
		exchangeName string
		pair         string
		jsonData	 []byte
		wantCode     int
	}{
		{
			name:         "exchange_name and pair are required",
			exchangeName: "",
			pair:         "",
			jsonData:	  []byte{},
			wantCode:     http.StatusBadRequest,
		},
		{
			name:         "exchange_name required",
			exchangeName: exchangeName,
			pair:         "",
			jsonData:	  []byte{},
			wantCode:     http.StatusBadRequest,
		},
		{	
			name:         "pair required",
			exchangeName: "",
			pair:         pair,
			jsonData:	  []byte{},
			wantCode:     http.StatusBadRequest,
		},
		{
			name:         "jsonData fail",
			exchangeName: exchangeName,
			pair:         pair,
			jsonData:	  []byte{},
			wantCode:     http.StatusBadRequest,
		},
		{
			name:         "jsonData",
			exchangeName: exchangeName,
			pair:         pair,
			jsonData:	  jsonData,
			wantCode:     http.StatusAccepted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewHandlers(&service.MockServiceMethods{})
			reqBody := strings.NewReader(string(tt.jsonData))
			req := httptest.NewRequest("POST", "/orderbook?exchange_name="+tt.exchangeName+"&pair="+tt.pair, reqBody)

			w := httptest.NewRecorder()
			handler.SaveOrderBookHandler()(w, req)

			if w.Code != tt.wantCode {
				t.Errorf("handler.SaveOrderBookHandler() = %v, want %v", w.Code, tt.wantCode)
			}
		})
	}
}

// тест для эндпоинта GET /order
func TestGetOrderHistoryHandler(t *testing.T) {
	client := models.Client{
		ClientName: "",
		ExchangeName: "test2",
		Label: "test3",
		Pair: "test4",
	}
	jsonData, err := json.Marshal(client)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name         string
		jsonData	 []byte
		wantCode     int
	}{
		{
			name:         "jsonData fail",
			jsonData:	  []byte{},
			wantCode:     http.StatusBadRequest,
		},
		{
			name:         "jsonData",
			jsonData:	  jsonData,
			wantCode:     http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewHandlers(&service.MockServiceMethods{})
			reqBody := strings.NewReader(string(tt.jsonData))
			req := httptest.NewRequest("POST", "/orderhistory", reqBody)

			w := httptest.NewRecorder()
			handler.GetOrderHistoryHandler()(w, req)
			if w.Code != tt.wantCode {
				t.Errorf("handler.GetOrderHistoryHandler() = %v, want %v", w.Code, tt.wantCode)
			}
		})
	}
}

// тест для эндпоинта POST /order
func TestSaveOrderHandler(t *testing.T) {
	order := models.HistoryOrder{
		ClientName:   "test1",
		ExchangeName: "test2",
		Label:		 "test3",
		Pair:		 "test4",
	}
	jsonData, err := json.Marshal(order)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name         string
		jsonData	 []byte
		wantCode     int
	}{
		{
			name:         "jsonData fail",
			jsonData:	  []byte{},
			wantCode:     http.StatusBadRequest,
		},
		{
			name:         "jsonData",
			jsonData:	  jsonData,
			wantCode:     http.StatusAccepted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewHandlers(&service.MockServiceMethods{})
			reqBody := strings.NewReader(string(tt.jsonData))
			req := httptest.NewRequest("POST", "/order", reqBody)

			w := httptest.NewRecorder()
			handler.SaveOrderHandler()(w, req)
			if w.Code != tt.wantCode {
				t.Errorf("handler.SaveOrderHandler() = %v, want %v", w.Code, tt.wantCode)
			}
		})
	}
}
