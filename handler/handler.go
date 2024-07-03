package handler

import (
	"context"

	"SynaPedia/usecase"
)

// Usecase represents interface of usecase that needed in handler
type Usecase interface {
	GetProductList(ctx context.Context, productCategoryID int64) ([]usecase.Product, error)

	Login(ctx context.Context, username string, password string) (usecase.LoginResponse, error)

	Register(ctx context.Context, param usecase.RegisterRequest) error

	AddToCart(ctx context.Context, param usecase.AddToCartRequest) error

	DeleteFromCart(ctx context.Context, cartID int64) error

	GetCartList(ctx context.Context, userID int64) ([]usecase.Cart, error)

	CreateOrder(ctx context.Context, param usecase.CreateOrderRequest) error
}

type Handler struct {
	Usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}
