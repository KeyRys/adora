package domain

type Product struct {
	ID       string  `json:"id"`
	SellerID string  `json:"seller_id"`
	Name     string  `json:"name"`
	Breed    string  `json:"breed"`
	Weight   float64 `json:"weight"`
	Color    string  `json:"color"`
	Gender   string  `json:"gender"`
	Price    float64 `json:"price"`
}
