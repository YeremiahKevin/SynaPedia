package handler

import (
	"context"

	"SynaPedia/usecase"
)

// ProductUsecase represents interface of product that needed in usecase
type ProductUsecase interface {
	GetProductList(ctx context.Context, productCategoryID int64) ([]usecase.Product, error)
}

type Handler struct {
	Product ProductUsecase
}

func NewHandler(product ProductUsecase) *Handler {
	return &Handler{
		Product: product,
	}
}
