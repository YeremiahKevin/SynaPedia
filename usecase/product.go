package usecase

import (
	"context"
	"log"
)

func (usecase *Usecase) GetProductList(ctx context.Context, productCategoryID int64) ([]Product, error) {
	productList, err := usecase.Product.GetProductList(ctx, productCategoryID)
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
