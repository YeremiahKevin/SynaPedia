package handler

import (
	"SynaPedia/usecase"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (handler *Handler) AddToCart(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body AddToCartRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Println("Handler AddToCart json.NewDecoder err: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctxHandler, _ := context.WithTimeout(context.Background(), time.Duration(time.Minute*1))

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

func (handler *Handler) DeleteFromCart(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cartIDStr := req.URL.Query().Get("cart_id")
	var cartID int64
	if cartIDStr != "" {
		cartID, _ = strconv.ParseInt(cartIDStr, 10, 64)
	}

	ctxHandler, _ := context.WithTimeout(context.Background(), time.Duration(time.Minute*1))

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

func (handler *Handler) GetCartList(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := req.URL.Query().Get("user_id")
	var userID int64
	if userIDStr != "" {
		userID, _ = strconv.ParseInt(userIDStr, 10, 64)
	}

	ctxHandler, _ := context.WithTimeout(context.Background(), time.Duration(time.Minute*1))

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
