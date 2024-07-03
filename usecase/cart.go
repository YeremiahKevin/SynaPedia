package usecase

import (
	"SynaPedia/repository"
	"context"
	"log"
)

func (usecase *Usecase) AddToCart(ctx context.Context, param AddToCartRequest) error {
	err := usecase.Repository.AddToCart(ctx, repository.AddToCartRequest(param))
	if err != nil {
		log.Println("Usecase AddToCart err: ", err)
		return err
	}

	return nil
}

func (usecase *Usecase) DeleteFromCart(ctx context.Context, cartID int64) error {
	err := usecase.Repository.DeleteFromCart(ctx, cartID)
	if err != nil {
		log.Println("Usecase DeleteFromCart err: ", err)
		return err
	}

	return nil
}

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
