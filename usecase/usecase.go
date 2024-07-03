package usecase

import (
	"context"

	"SynaPedia/repository"
)

// ProductRepository represents interface of product that needed in usecase
type ProductRepository interface {
	GetProductList(ctx context.Context, productCategoryID int64) ([]repository.Product, error)
}

type Usecase struct {
	Product ProductRepository
}

func NewUsecase(product ProductRepository) *Usecase {
	return &Usecase{
		Product: product,
	}
}
