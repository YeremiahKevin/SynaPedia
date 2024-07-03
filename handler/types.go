package handler

type GeneralResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type AddToCartRequest struct {
	UserID       int64 `json:"user_id"`
	ProductSkuID int64 `json:"product_sku_id"`
	Quantity     int64 `json:"quantity"`
}
