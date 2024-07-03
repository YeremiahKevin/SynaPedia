package repository

type Product struct {
	ProductID         int64  `db:"product_id"`
	ProductCategoryID int64  `db:"product_category_id"`
	ProductName       string `db:"product_name"`
	Price             int64  `db:"price"`
	Description       string `db:"description"`
}
