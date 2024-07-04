package handler

import (
	// golang package
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	// internal package
	"SynaPedia/usecase"
)

// AddToCart is a function to handle add to cart request
// it accepts http.ResponseWriter and pointer of http.Request as parameters
// it returns status code 200 when success
// otherwise it returns status code 400 when request invalid or status code 500 when error occurs
func (handler *Handler) AddToCart(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctxHandler, _ := context.WithTimeout(context.Background(), time.Duration(time.Minute*1))

	var body AddToCartRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Println("Handler AddToCart json.NewDecoder err: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.Usecase.AddToCart(ctxHandler, usecase.AddToCartRequest(body))
	if err != nil {
		log.Println("Handler AddToCart handler.Usecase.AddToCart err: ", err)
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

// DeleteFromCart is a function to handle delete from cart request
// it accepts http.ResponseWriter and pointer of http.Request as parameters
// it returns status code 200 when success
// otherwise it returns status code 400 when request invalid or status code 500 when error occurs
func (handler *Handler) DeleteFromCart(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctxHandler, _ := context.WithTimeout(context.Background(), time.Duration(time.Minute*1))

	cartIDStr := req.URL.Query().Get("cart_id")
	var cartID int64
	if cartIDStr != "" {
		cartID, _ = strconv.ParseInt(cartIDStr, 10, 64)
	}
	if cartID <= 0 {
		log.Println("Handler DeleteFromCart cartID <= 0")
		http.Error(w, "Invalid cart id", http.StatusBadRequest)
		return
	}

	err := handler.Usecase.DeleteFromCart(ctxHandler, cartID)
	if err != nil {
		log.Println("Handler DeleteFromCart handler.Usecase.DeleteFromCart err: ", err)
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

// GetCartList is a function to get cart list request
// it accepts http.ResponseWriter and pointer of http.Request as parameters
// it returns status code 200 when success
// otherwise it returns status code 400 when request invalid or status code 500 when error occurs
func (handler *Handler) GetCartList(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctxHandler, _ := context.WithTimeout(context.Background(), time.Duration(time.Minute*1))

	userIDStr := req.URL.Query().Get("user_id")
	var userID int64
	if userIDStr != "" {
		userID, _ = strconv.ParseInt(userIDStr, 10, 64)
	}
	if userID <= 0 {
		log.Println("Handler GetCartList userID <= 0")
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	result, err := handler.Usecase.GetCartList(ctxHandler, userID)
	if err != nil {
		log.Println("Handler GetCartList handler.Usecase.GetCartList err: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := GeneralResponse{
		Status: "Success",
		Data:   result,
	}
	_ = json.NewEncoder(w).Encode(response)

	return
}
