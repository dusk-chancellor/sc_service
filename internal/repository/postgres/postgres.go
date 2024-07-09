package postgres

// подключение к бд и создание таблиц

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // драйвер постгреса
	"github.com/dusk-chancellor/sc_service/configs"
)

type Storage struct {
	db *sql.DB
	// ...
}
// все данные для подключение запрашиваются из конфига
func ConnectDB(cfg *configs.Config) (*Storage, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
			cfg.Database.DbUser,
			cfg.Database.DbPassword,
			cfg.Database.DbName,
			cfg.Database.DbHost,
			cfg.Database.DbPort,
		),
	)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err = createTables(db); err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func createTables(db *sql.DB) error {
	// asks & bids сделаны в виде жсон
	const (
		tablesQuery = `
			CREATE TABLE IF NOT EXISTS order_book (
    			id BIGSERIAL PRIMARY KEY,
    			exchange VARCHAR(255) NOT NULL,
    			pair VARCHAR(255) NOT NULL,
    			asks JSONB,
    			bids JSONB
			);


			CREATE TABLE IF NOT EXISTS order_history (
				client_name VARCHAR(255) NOT NULL,
				exchange VARCHAR(255) NOT NULL,
				label VARCHAR(255) NOT NULL,
				pair VARCHAR(255) NOT NULL,
				side VARCHAR(255) NOT NULL,
				type VARCHAR(255) NOT NULL,
				base_qty FLOAT NOT NULL,
				price FLOAT NOT NULL,
				algorithm_name_placed VARCHAR(255),
				lowest_sell_prc FLOAT,
				highest_buy_prc FLOAT,
				commission_quote_qty FLOAT,
				time_placed TIMESTAMP NOT NULL DEFAULT NOW()
			);
		`
	)

	if _, err := db.Exec(tablesQuery); err != nil {
		return err
	}
	return nil
}
