package domain

type Order struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	TotalPrice int    `json:"total_price"`
	Status     string `json:"status"`
}

type OrderItem struct {
	ID       string `json:"id"`
	OrderID  string `json:"order_id"`
	RabbitID string `json:"rabbit_id"`
	SellerID string `json:"seller_id"`
	Price    int    `json:"price"`
}

type CheckoutItem struct {
	RabbitID string `json:"rabbit_id"`
	SellerID string `json:"seller_id"`
	Price    int    `json:"price"`
}
