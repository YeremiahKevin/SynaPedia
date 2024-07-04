package handler

import (
	// golang package
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	// external package
	"github.com/kataras/jwt"

	// internal package
	"SynaPedia/usecase"
)

// Login is a function to handle user login request
// it accepts http.ResponseWriter and pointer of http.Request as parameters
// it returns status code 200 when success
// otherwise it returns status code 400 when request invalid or status code 500 when error occurs
func (handler *Handler) Login(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctxHandler, _ := context.WithTimeout(context.Background(), time.Duration(time.Minute*1))

	var body LoginRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Println("Handler Login json.NewDecoder err: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := handler.Usecase.Login(ctxHandler, body.Username, body.Password)
	if err != nil {
		log.Println("Handler Login handler.Usecase.Login err: ", err)
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

// Register is a function to handle register new user request
// it accepts http.ResponseWriter and pointer of http.Request as parameters
// it returns status code 200 when success
// otherwise it returns status code 400 when request invalid or status code 500 when error occurs
func (handler *Handler) Register(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		log.Println("Handler Register req.Method err: " + req.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctxHandler, _ := context.WithTimeout(context.Background(), time.Duration(time.Minute*1))

	var body RegisterRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Println("Handler Register json.NewDecoder err: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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
