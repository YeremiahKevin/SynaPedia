package repository

import (
	// golang package
	"context"
	"fmt"
	"log"
)

// GetProductList is a function to get product list from product
// it accepts context.Context and int64 as parameters
// it returns non-empty slice of Product and nil error when success
// otherwise it returns nil and detailed error
func (repository *Repository) GetProductList(ctx context.Context, productCategoryID int64) ([]Product, error) {
	var productList []Product

	if productCategoryID == 0 {
		queryResult, err := db.QueryContext(ctx, "SELECT product_id, product_category_id, product_name, description FROM product")
		if err != nil {
			log.Println("Repository GetProductList db.QueryContext err: ", err)
			return nil, fmt.Errorf("failed to get product list: %s", err.Error())
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
		queryResult, err := db.QueryContext(ctx, "SELECT product_id, product_category_id, product_name, description FROM product WHERE product_category_id = $1", productCategoryID)
		if err != nil {
			log.Println("Repository GetProductList db.QueryContext err: ", err)
			return nil, fmt.Errorf("failed to get product list: %s", err.Error())
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
