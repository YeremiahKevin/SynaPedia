package repository

import (
	"context"
	"fmt"
	"log"
	"time"
)

func (repository *Repository) AddToCart(ctx context.Context, param AddToCartRequest) error {
	_, err := db.ExecContext(ctx, `INSERT INTO cart(user_id, product_sku_id, quantity, created_at) VALUES($1, $2, $3, $4)`, param.UserID, param.ProductSkuID, param.Quantity, time.Now())
	if err != nil {
		log.Println("Repository AddToCart db.ExecContext err: ", err)
		return fmt.Errorf("failed to add to cart: %s", err.Error())
	}

	return nil
}

func (repository *Repository) DeleteFromCart(ctx context.Context, cartID int64) error {
	_, err := db.ExecContext(ctx, `UPDATE cart
										SET deleted_at = $1
										WHERE cart_id = $2`, time.Now(), cartID)
	if err != nil {
		log.Println("Repository DeleteFromCart db.ExecContext err: ", err)
		return fmt.Errorf("failed to delete from cart: %s", err.Error())
	}

	return nil
}

func (repository *Repository) GetCartList(ctx context.Context, userID int64) ([]Cart, error) {
	var cartList []Cart

	queryResult, err := db.QueryContext(ctx, `SELECT cart.cart_id,
														   cart.quantity,
														   product.product_id,
														   product.product_name,
														   product_sku.product_sku_id,
														   product_sku.variant,
														   product_sku.price,
														   product_sku.stock
													FROM cart
															 JOIN product_sku
																  ON product_sku.product_sku_id = cart.product_sku_id
																	  AND product_sku.deleted_at ISNULL
															 JOIN product
																  ON product.product_id = product_sku.product_id
																	  AND product.deleted_at ISNULL
													WHERE cart.user_id = $1
													  AND cart.deleted_at ISNULL`, userID)
	if err != nil {
		log.Println("Repository GetCartList db.Query err: ", err)
		return nil, fmt.Errorf("failed to get cart list data: %s", err.Error())
	}

	for queryResult.Next() {
		var cart Cart
		err = queryResult.Scan(&cart.CartID, &cart.Quantity, &cart.ProductID, &cart.ProductName, &cart.ProductSkuID, &cart.Variant, &cart.Price, &cart.Stock)
		if err != nil {
			log.Println("Repository GetCartList queryResult.Scan err: ", err)
		}
		cartList = append(cartList, cart)
	}

	return cartList, nil
}
