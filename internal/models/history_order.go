package models

import "time"

// модельная структура для таблицы ордер хистори
type HistoryOrder struct {
	ClientName          string `json:"client_name"`
	ExchangeName   	  	string `json:"exchange_name"`
	Label		        string `json:"label"`
	Pair  		        string `json:"pair"`
	Side    		    string `json:"side,omitempty"`
	Type            	string `json:"type,omitempty"`
	BaseQty             float64 `json:"base_qty,omitempty"`
	Price               float64 `json:"price,omitempty"`
	AlgorithmNamePlaced string `json:"algorithm_name_placed,omitempty"`
	LowestSellPrc       float64 `json:"lowest_sell_prc,omitempty"`
	HighestBuyPrc       float64 `json:"highest_buy_prc,omitempty"`
	CommissionQuoteQty  float64 `json:"commission_quote_qty,omitempty"`
	TimePlaced          time.Time // будет автоматом устанавливаться в самой бд
}
	