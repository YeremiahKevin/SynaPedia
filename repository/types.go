package repository

type Product struct {
	ProductID         int64  `db:"product_id"`
	ProductCategoryID int64  `db:"product_category_id"`
	ProductName       string `db:"product_name"`
	Price             int64  `db:"price"`
	Description       string `db:"description"`
}

type AddToCartRequest struct {
	UserID       int64 `db:"user_id"`
	ProductSkuID int64 `db:"product_id"`
	Quantity     int64 `db:"quantity"`
}

type Cart struct {
	CartID       int64  `db:"cart_id"`
	Quantity     int64  `db:"quantity"`
	ProductID    int64  `db:"product_id"`
	ProductName  string `db:"product_name"`
	ProductSkuID string `db:"product_sku_id"`
	Variant      string `db:"variant"`
	Price        int64  `db:"price"`
	Stock        int64  `db:"stock"`
}
