package handler

import (
	// golang package
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	// internal package
	"SynaPedia/usecase"
)

// CreateOrder is a function to handle create new order request
// it accepts http.ResponseWriter and pointer of http.Request as parameters
// it returns status code 200 when success
// otherwise it returns status code 400 when request invalid or status code 500 when error occurs
func (handler *Handler) CreateOrder(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctxHandler, _ := context.WithTimeout(context.Background(), time.Duration(time.Minute*1))

	var body CreateOrderRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Println("Handler CreateOrder json.NewDecoder err: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderDetails := make([]usecase.OrderDetail, len(body.OrderDetails))
	for i, detail := range body.OrderDetails {
		orderDetails[i] = usecase.OrderDetail(detail)
	}
	err = handler.Usecase.CreateOrder(ctxHandler, usecase.CreateOrderRequest{
		UserID:        body.UserID,
		UserAddressID: body.UserAddressID,
		OrderDetails:  orderDetails,
		PaymentID:     body.PaymentID,
		Amount:        body.Amount,
	})
	if err != nil {
		log.Println("Handler CreateOrder handler.Usecase.CreateOrder err: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := GeneralResponse{
		Status: "Success",
		Data:   nil,
	}
	_ = json.NewEncoder(w).Encode(response)

	return
}
