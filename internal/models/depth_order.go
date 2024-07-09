package models

// модельная структура для хранения ордер
type DepthOrder struct {
	Price 	float64 `json:"price"`
	BaseQty float64 `json:"base_qty"`
}