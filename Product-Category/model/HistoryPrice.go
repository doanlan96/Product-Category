package model

type HistoryPrice struct {
	Id        int64   `json:"id"`
	ProductId int64   `json:"productId"`
	OldPrice  float64 `json:"oldPrice"`
}
