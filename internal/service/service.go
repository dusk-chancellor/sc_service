package service

// сервисные методы, выполняют основную бизнес логику приложения
// берут данные с бд и отправляют хендлерам

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dusk-chancellor/sc_service/internal/models"
)

type Service struct {
	DbMethods DbMethods
	Ctx context.Context // передадим бд, по тз в сервисных методах они не применяются
	// ...
}
// взаимодействие с бд в интерфейсе
type DbMethods interface {
	CreateOrder(ctx context.Context, client *models.Client, order *models.HistoryOrder) error
	CreateOrderBook(ctx context.Context, exchange_name, pair string, asks, bids []byte) error
	FetchOrderHistory(ctx context.Context, client *models.Client) ([]*models.HistoryOrder, error)
	FetchOrderBook(ctx context.Context, exchange_name, pair string) ([]byte, []byte, error)
}
// инициализация в ~/internal/app/app.go
func NewService(db DbMethods, ctx context.Context) *Service {
	return &Service{
		DbMethods: db,
		Ctx:       ctx,
	}
}

func (s *Service) SaveOrder(client *models.Client, order *models.HistoryOrder) error {
	err := s.DbMethods.CreateOrder(s.Ctx, client, order)
	if err != nil {
		log.Printf("error saving order: %v", err)
		return err
	}
	return nil
}

func (s *Service) GetOrderHistory(client *models.Client)  ([]*models.HistoryOrder, error) {
	
	orderHistory, err := s.DbMethods.FetchOrderHistory(s.Ctx, client)
	if err != nil {
		log.Printf("error fetching order history: %v", err)
		return nil, err
	}
	return orderHistory, nil
}

func (s *Service) SaveOrderBook(exchange_name, pair string, orderBook []*models.DepthOrder) error {
	// не ясно как разделять ордербук на asks и bids для сохранения в бд
	// поэтому решил просто: asks - положительные числа, bids - отрицательные
	var asks, bids []*models.DepthOrder
	for _, order := range orderBook {
		if order.BaseQty > 0 {
			asks = append(asks, order)
		} else {
			bids = append(bids, order)
		}
	}
	// превращаем в жсон байты
	asksJSON, err := json.Marshal(asks)
	if err != nil {
		log.Printf("error marshaling asks: %v", err)
		return err
	}

	bidsJSON, err := json.Marshal(bids)
	if err != nil {
		log.Printf("error marshaling bids: %v", err)
		return err
	}

	err = s.DbMethods.CreateOrderBook(s.Ctx, exchange_name, pair, asksJSON, bidsJSON)
	if err != nil {
		log.Printf("error saving orderbook: %v", err)
		return err
	}

	return nil
}

func (s *Service) GetOrderBook(exchange_name, pair string) ([]*models.DepthOrder, error) {
	asks, bids , err := s.DbMethods.FetchOrderBook(s.Ctx, exchange_name, pair)
	if err != nil {
		log.Printf("error fetching orderbook: %v", err)
		return nil, err
	}

	var asksJSON, bidsJSON []*models.DepthOrder
	if err = json.Unmarshal(asks, &asksJSON); err != nil {
		log.Printf("error unmarshaling asks: %v", err)
		return nil, err
	}
	if err = json.Unmarshal(bids, &bidsJSON); err != nil {
		log.Printf("error unmarshaling bids: %v", err)
		return nil, err
	}

	var orderBook []*models.DepthOrder
	orderBook = append(orderBook, asksJSON...)
	orderBook = append(orderBook, bidsJSON...)
	return orderBook, nil
}
