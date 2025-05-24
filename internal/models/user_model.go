package types

type User struct {
	ID           int     `json:"id"`
	Name         string  `json:"name" validate:"required" min=2,max=30`
	Email        string  `json:"email" validate:"required,email"`
	Phone        string  `json:"phone" validate:"required"`
	Password     string  `json:"password" validate:"required,min=6"`
	CreatedAt    string  `json:"created_at"`
	Token        string  `json:"token"`
	RefreshToken string  `json:"refresh_token"`
	UpdatedAt    string  `json:"updated_at"`
	DeletedAt    string  `json:"deleted_at"`
	Address      Address `json:"address" validate:"required"`
	Cart         []Cart  `json:"cart"`
	Orders       []Order `json:"orders"`
}

type Address struct {
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	ZipCode   string `json:"zip_code"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
