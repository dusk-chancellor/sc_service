package models

// модельная структура для данных о клиенте
type Client struct{
	ClientName   string `json:"client_name"`
	ExchangeName string `json:"exchange_name"`
	Label		 string `json:"label"`
	Pair  		 string `json:"pair"`
}
	