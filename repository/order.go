package repository

import (
	"context"
	"fmt"
	"log"
	"time"
)

func (repository *Repository) CreateOrder(ctx context.Context, param CreateOrderRequest) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Repository CreateOrder db.BeginTx err: ", err)
		return fmt.Errorf("failed to begin transaction: %s", err.Error())
	}

	defer tx.Rollback()

	resultOrder, err := tx.ExecContext(ctx, `INSERT INTO "order"(user_id, user_address_id, total_price, created_at) VALUES($1, $2, $3, $4)`, param.UserID, param.UserAddressID, param.TotalPrice, time.Now())
	if err != nil {
		log.Println("Repository CreateOrder db.ExecContext err: ", err)
		return fmt.Errorf("failed to create order: %s", err.Error())
	}

	orderID, err := resultOrder.LastInsertId()
	if err != nil {
		log.Println("Repository CreateOrder resultOrder.LastInsertId err: ", err)
		return fmt.Errorf("failed to get order id: %s", err.Error())
	}

	for i := 0; i < len(param.OrderDetails); i++ {
		_, err = tx.ExecContext(ctx, `INSERT INTO order_detail(order_id, product_sku_id, quantity, price, created_at) VALUES($1, $2, $3, $4, $5)`, orderID, param.OrderDetails[i].ProductSkuID, param.OrderDetails[i].Quantity, param.OrderDetails[i].Price, time.Now())
		if err != nil {
			log.Println("Repository CreateOrder db.ExecContext err: ", err)
			return fmt.Errorf("failed to create order detail: %s", err.Error())
		}
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO order_payment(order_id, payment_id, amount, status, created_at) VALUES($1, $2, $3, $4, $5)`, orderID, param.PaymentID, param.Amount, param.Status, time.Now())
	if err != nil {
		log.Println("Repository CreateOrder db.ExecContext err: ", err)
		return fmt.Errorf("failed to create order payment: %s", err.Error())
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Repository CreateOrder tx.Commit err: ", err)
		return fmt.Errorf("failed to commit: %s", err.Error())
	}

	return nil
}
