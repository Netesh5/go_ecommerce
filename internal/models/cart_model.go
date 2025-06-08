package models

type Cart struct {
	ID        int    `json:"id" validate:"required"`
	UserID    int    `json:"user_id" validate:"required"`
	ProductID int    `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
	Price     string `json:"price" validate:"required"`
	Total     string `json:"total"`
	Checkout  bool   `json:"checkout"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type Order struct {
	ID          int     `json:"id" validate:"required"`
	UserID      int     `json:"user_id" validate:"required"`
	ProductID   int     `json:"product_id" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	TotalPrice  float64 `json:"total_price"`
	OrderStatus string  `json:"order_status"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   string  `json:"deleted_at"`
}

type UpdateCartReq struct {
	Id       int `json:"id" validate:"required"`
	Product  int `json:"product_id" validate:"required"`
	Quantity int `json:"quantity" validate:"required"`
}
