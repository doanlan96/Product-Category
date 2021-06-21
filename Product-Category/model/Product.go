package model

type Product struct {
	Id         int64    `json:"id"`
	CategoryID int64    `json:"categoryId"`
	Images     []string `json:"image"`
	Name       string   `json:"name"`
	Price      float64  `json:"price"`
	IsSale     bool     `json:"isSale"`
	// Reviews       []Review `json:"reviews"`
	PriceHistory  []float64 `json:"priceHistory"`
	AverageRating float32   `json:"averageRating"`
	CreatedAt     int64     `json:"createdAt"`
	ModifiedAt    int64     `json:"modifiedAt"`
}
