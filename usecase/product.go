package usecase

import (
	// golang package
	"context"
	"log"
)

// GetProductList is a function to get list of product
// it accepts context.Context and int64 as parameters
// it returns non-empty slice of Product and nil error when success
// otherwise it returns nil and detailed error
func (usecase *Usecase) GetProductList(ctx context.Context, productCategoryID int64) ([]Product, error) {
	productList, err := usecase.Repository.GetProductList(ctx, productCategoryID)
	if err != nil {
		log.Println("Usecase GetProductList err: ", err)
		return nil, err
	}

	result := make([]Product, len(productList))
	for i, product := range productList {
		result[i] = Product(product)
	}

	return result, nil
}
