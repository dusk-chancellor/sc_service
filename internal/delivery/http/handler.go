package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dusk-chancellor/sc_service/internal/models"
)

type Handlers struct {
	ServiceMethods ServiceMethods
}

// методы сервиса, реализованные в интерфейсе
// основные функции с тз
type ServiceMethods interface {
	GetOrderBook(exchange_name, pair string) ([]*models.DepthOrder, error)
	SaveOrderBook(exchange_name, pair string, orderBook []*models.DepthOrder) error
	GetOrderHistory(client *models.Client)  ([]*models.HistoryOrder, error)
	SaveOrder(client *models.Client, order *models.HistoryOrder) error
}

func NewHandlers(serviceMethods ServiceMethods) *Handlers {
	return &Handlers{
		ServiceMethods: serviceMethods,
	}
}

// Эндпоинт: GET /orderbook
func (h *Handlers) GetOrderBookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// exchange_name и pair получаем с параметров запроса
		exchange_name := r.URL.Query().Get("exchange_name")
		pair := r.URL.Query().Get("pair")
		if exchange_name == "" || pair == "" {
			http.Error(w, "exchange_name and pair are required", http.StatusBadRequest)
			return
		}

		orderBook, err := h.ServiceMethods.GetOrderBook(exchange_name, pair)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// возвращаем ордер в боди ответа (жсон)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(orderBook); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("successful op: GetOrderBookHandler")
		// Жсон энкодер сам вышлет статус в хедере
	}
}

// Эндпоинт: POST /orderbook
func (h *Handlers) SaveOrderBookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		exchange_name := r.URL.Query().Get("exchange_name")
		pair := r.URL.Query().Get("pair")
		if exchange_name == "" || pair == "" {
			http.Error(w, "exchange_name and pair are required", http.StatusBadRequest)
			return
		}

		// ордербук ожидается в боди запроса
		var orderBook []*models.DepthOrder
		if err := json.NewDecoder(r.Body).Decode(&orderBook); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.ServiceMethods.SaveOrderBook(exchange_name, pair, orderBook); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		log.Printf("successful op: SaveOrderBookHandler")
	}
}

// Эндпоинт: GET /order
func (h *Handlers) GetOrderHistoryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var client models.Client
		// читаем данные о клиенте с боди запроса (жсон)
		if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		orderHistory, err := h.ServiceMethods.GetOrderHistory(&client)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// возвращаем тоже в боди ответа
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(orderHistory); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("successful op: GetOrderHistoryHandler")
	}
}

// Эндпоинт: POST /order
func (h *Handlers) SaveOrderHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order models.HistoryOrder
		// читаем ордер с боди запроса
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		client := models.Client{
			ClientName:   order.ClientName,
			ExchangeName: order.ExchangeName,
			Label:		 order.Label,
			Pair:		 order.Pair,
		}
		if err := h.ServiceMethods.SaveOrder(&client, &order); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		log.Printf("successful op: SaveOrderHandler")
	}
}
