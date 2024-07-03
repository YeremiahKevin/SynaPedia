package repository

import (
	"context"
	"log"
)

func (repository *Repository) GetProductList(ctx context.Context, productCategoryID int64) ([]Product, error) {
	var productList []Product

	if productCategoryID == 0 {
		queryResult, err := db.QueryContext(ctx, "SELECT product_id, product_category_id, product_name, description FROM product")
		if err != nil {
			log.Println("Repository GetProductList db.Query err: ", err)
		}

		for queryResult.Next() {
			var product Product
			err = queryResult.Scan(&product.ProductID, &product.ProductCategoryID, &product.ProductName, &product.Description)
			if err != nil {
				log.Println("Repository GetProductList queryResult.Scan err: ", err)
			}
			productList = append(productList, product)
		}
	} else {
		queryResult, err := db.Query("SELECT product_id, product_category_id, product_name, description FROM product WHERE product_category_id = $1", productCategoryID)
		if err != nil {
			log.Println("Repository GetProductList db.Query err: ", err)
		}

		for queryResult.Next() {
			var product Product
			err = queryResult.Scan(&product.ProductID, &product.ProductCategoryID, &product.ProductName, &product.Description)
			if err != nil {
				log.Println("Repository GetProductList queryResult.Scan err: ", err)
			}
			productList = append(productList, product)
		}
	}

	return productList, nil
}
