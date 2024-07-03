package usecase

type Product struct {
	ProductID         int64  `json:"product_id"`
	ProductCategoryID int64  `json:"product_category_id"`
	ProductName       string `json:"product_name"`
	Price             int64  `json:"price"`
	Description       string `json:"description"`
}

type LoginResponse struct {
	IsAllowed bool
}

type RegisterRequest struct {
	Username string
	Password string
	FullName string
}
