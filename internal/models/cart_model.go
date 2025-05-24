package models

type Cart struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     string `json:"price"`
	Total     string `json:"total"`
	Checkout  bool   `json:"checkout"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type Order struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	ProductID   int     `json:"product_id"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
	OrderStatus string  `json:"order_status"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   string  `json:"deleted_at"`
}
