package usecase

import (
	// golang package
	"context"
	"log"

	// internal package
	"SynaPedia/repository"
)

// CreateOrder is a function to create new order
// it accepts context.Context and CreateOrderRequest as parameters
// it returns nil error when success
// otherwise it returns detailed error
func (usecase *Usecase) CreateOrder(ctx context.Context, param CreateOrderRequest) error {
	var totalPrice int64 = 0
	orderDetails := make([]repository.OrderDetail, len(param.OrderDetails))
	for i := 0; i < len(param.OrderDetails); i++ {
		totalPrice += param.OrderDetails[i].Price
		orderDetails[i] = repository.OrderDetail{
			ProductSkuID: param.OrderDetails[i].ProductSkuID,
			Quantity:     param.OrderDetails[i].Quantity,
			Price:        param.OrderDetails[i].Price,
		}
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
		log.Println("Usecase CreateOrder usecase.Repository.CreateOrder err: ", err)
		return err
	}

	for _, orderDetail := range param.OrderDetails {
		err = usecase.Repository.DeleteFromCart(ctx, orderDetail.CartID)
		if err != nil {
			log.Println("Usecase CreateOrder usecase.Repository.DeleteFromCart err: ", err)
		}
	}

	return nil
}
