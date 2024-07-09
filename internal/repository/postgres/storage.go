package postgres

// все методы для взаимодействий с бд

import (
	"context"

	"github.com/dusk-chancellor/sc_service/internal/models"
)

func (s *Storage) CreateOrderBook(ctx context.Context, exchange_name, pair string, asks, bids []byte) error {

	query, err := s.db.PrepareContext(ctx,
		"INSERT INTO order_book (exchange, pair, asks, bids) VALUES ($1, $2, $3::jsonb, $4::jsonb)",
	)
	if err != nil {
		return err
	}

	_, err = query.ExecContext(ctx, exchange_name, pair, asks, bids)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) FetchOrderBook(ctx context.Context, exchange_name, pair string) ([]byte, []byte, error) {

	query, err := s.db.PrepareContext(ctx,
		"SELECT asks, bids FROM order_book WHERE exchange = $1 AND pair = $2",
	)
	if err != nil {
		return nil, nil, err
	}

	var asks, bids []byte
	row := query.QueryRowContext(ctx, exchange_name, pair)
	if err = row.Scan(&asks, &bids); err != nil {
		return nil, nil, err
	}
	// возвращает жсон байты
	return asks, bids, nil
}

func (s *Storage) CreateOrder(ctx context.Context, client *models.Client, order *models.HistoryOrder) error {

	query, err := s.db.PrepareContext(ctx,
		"INSERT INTO order_history VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
	)
	if err != nil {
		return err
	}

	_, err = query.ExecContext(ctx,
		client.ClientName,   // не был уверен насчет того, что данные с ордер сюда подойдут
		client.ExchangeName, // было бы хорошо, если тз улучшили насчет этого всего
		client.Label,
		client.Pair,
		order.Side,
		order.Type,
		order.BaseQty,
		order.Price,
		order.AlgorithmNamePlaced,
		order.LowestSellPrc,
		order.HighestBuyPrc,
		order.CommissionQuoteQty,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) FetchOrderHistory(ctx context.Context, client *models.Client) ([]*models.HistoryOrder, error) {
	// ...
	query, err := s.db.PrepareContext(ctx,
		"SELECT * FROM order_history WHERE client_name = $1",
	)
	if err != nil {
		return nil, err
	}

	row, err := query.QueryContext(ctx, client.ClientName)
	if err != nil {
		return nil, err
	}

	var orderHistory []*models.HistoryOrder
	for row.Next() {
		var order models.HistoryOrder
		err = row.Scan(
			&order.ClientName,
			&order.ExchangeName,
			&order.Label,
			&order.Pair,
			&order.Side,
			&order.Type,
			&order.BaseQty,
			&order.Price,
			&order.AlgorithmNamePlaced,
			&order.LowestSellPrc,
			&order.HighestBuyPrc,
			&order.CommissionQuoteQty,
			&order.TimePlaced,
		)
		if err != nil {
			return nil, err
		}
		orderHistory = append(orderHistory, &order)
	}
	return orderHistory, nil
}
