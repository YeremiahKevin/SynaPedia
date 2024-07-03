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

type AddToCartRequest struct {
	UserID       int64
	ProductSkuID int64
	Quantity     int64
}

type Cart struct {
	CartID       int64  `json:"cart_id"`
	Quantity     int64  `json:"quantity"`
	ProductID    int64  `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductSkuID string `json:"product_sku_id"`
	Variant      string `json:"variant"`
	Price        int64  `json:"price"`
	Stock        int64  `json:"stock"`
}
