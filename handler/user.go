package handler

import (
	"SynaPedia/usecase"
	"context"
	"encoding/json"
	"github.com/kataras/jwt"
	"log"
	"net/http"
	"os"
	"time"
)

func (handler *Handler) Login(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body LoginRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctxHandler, _ := context.WithTimeout(context.Background(), time.Duration(time.Minute*1))

	result, err := handler.Usecase.Login(ctxHandler, body.Username, body.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var token string
	if result.IsAllowed {
		tokenByte, _ := jwt.Sign(jwt.HS256, []byte(os.Getenv("JWT_KEY")), map[string]interface{}{
			"username": body.Username,
		}, jwt.MaxAge(60*time.Minute))
		token = string(tokenByte)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := GeneralResponse{
		Status: "Success",
		Data: map[string]interface{}{
			"token":      token,
			"is_allowed": result.IsAllowed,
		},
	}
	_ = json.NewEncoder(w).Encode(response)

	return
}

func (handler *Handler) Register(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		log.Println("Handler Register req.Method err: " + req.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body RegisterRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Println("Handler Register json.NewDecoder err: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctxHandler, _ := context.WithTimeout(context.Background(), time.Duration(time.Minute*1))

	err = handler.Usecase.Register(ctxHandler, usecase.RegisterRequest(body))
	if err != nil {
		log.Println("Handler Register handler.Usecase.Register err: " + err.Error())
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
