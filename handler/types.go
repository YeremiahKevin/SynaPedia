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

type CreateOrderRequest struct {
	UserID        int64         `json:"user_id"`
	UserAddressID int64         `json:"user_address_id"`
	OrderDetails  []OrderDetail `json:"order_details"`
	PaymentID     int64         `json:"payment_id"`
	Amount        int64         `json:"amount"`
}

type OrderDetail struct {
	CartID       int64 `json:"cart_id"`
	ProductSkuID int64 `json:"product_sku_id"`
	Quantity     int64 `json:"quantity"`
	Price        int64 `json:"price"`
}
