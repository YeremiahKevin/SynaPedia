package usecase

import (
	"context"

	"SynaPedia/repository"
)

// Repository represents interface of repository that needed in usecase
type Repository interface {
	GetProductList(ctx context.Context, productCategoryID int64) ([]repository.Product, error)

	Login(ctx context.Context, username string, password string) (repository.LoginResponse, error)

	Register(ctx context.Context, param repository.RegisterRequest) error

	AddToCart(ctx context.Context, param repository.AddToCartRequest) error

	DeleteFromCart(ctx context.Context, cartID int64) error

	GetCartList(ctx context.Context, userID int64) ([]repository.Cart, error)

	CreateOrder(ctx context.Context, param repository.CreateOrderRequest) error
}

type Usecase struct {
	Repository Repository
}

func NewUsecase(repository Repository) *Usecase {
	return &Usecase{
		Repository: repository,
	}
}
