package usecase

import (
	"SynaPedia/repository"
	"context"
	"log"
)

func (usecase *Usecase) CreateOrder(ctx context.Context, param CreateOrderRequest) error {
	var totalPrice int64 = 0
	orderDetails := make([]repository.OrderDetail, len(param.OrderDetails))
	for i := 0; i < len(param.OrderDetails); i++ {
		totalPrice += param.OrderDetails[i].Price
		orderDetails[i] = repository.OrderDetail(param.OrderDetails[i])
	}

	err := usecase.Repository.CreateOrder(ctx, repository.CreateOrderRequest{
		UserID:        param.UserID,
		UserAddressID: param.UserAddressID,
		TotalPrice:    totalPrice,
		OrderDetails:  orderDetails,
		PaymentID:     param.PaymentID,
		Amount:        param.Amount,
		Status:        1, // waiting for payment,
	})
	if err != nil {
		log.Println("Usecase CreateOrder err: ", err)
		return err
	}

	return nil
}
