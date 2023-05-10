package domain

type Item struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Picture string  `json:"picture"`
	Link    string  `json:"link"`
}
