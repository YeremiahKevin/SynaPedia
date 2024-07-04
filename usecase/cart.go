package usecase

import (
	// golang package
	"context"
	"log"

	// internal package
	"SynaPedia/repository"
)

// AddToCart is a function to add new product to a cart
// it accepts context.Context and AddToCartRequest as parameters
// it returns nil error when success
// otherwise it returns detailed error
func (usecase *Usecase) AddToCart(ctx context.Context, param AddToCartRequest) error {
	err := usecase.Repository.AddToCart(ctx, repository.AddToCartRequest(param))
	if err != nil {
		log.Println("Usecase AddToCart err: ", err)
		return err
	}

	return nil
}

// DeleteFromCart is a function to delete existing product in a cart
// it accepts context.Context and int64 as parameters
// it returns nil error when success
// otherwise it returns detailed error
func (usecase *Usecase) DeleteFromCart(ctx context.Context, cartID int64) error {
	err := usecase.Repository.DeleteFromCart(ctx, cartID)
	if err != nil {
		log.Println("Usecase DeleteFromCart err: ", err)
		return err
	}

	return nil
}

// GetCartList is a function to get list of product in a cart
// it accepts context.Context and int64 as parameters
// it returns non-empty slice of Cart and nil error when success
// otherwise it returns nil and detailed error
func (usecase *Usecase) GetCartList(ctx context.Context, userID int64) ([]Cart, error) {
	cartList, err := usecase.Repository.GetCartList(ctx, userID)
	if err != nil {
		log.Println("Usecase GetCartList err: ", err)
		return nil, err
	}

	result := make([]Cart, len(cartList))
	for i, cart := range cartList {
		result[i] = Cart(cart)
	}

	return result, nil
}
