package main

import (
	"SynaPedia/handler"
	"SynaPedia/repository"
	"SynaPedia/usecase"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	log.Println("Running handler...")

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	repo := repository.NewRepository()
	uc := usecase.NewUsecase(repo)
	serv := handler.NewHandler(uc)

	http.HandleFunc("/product-list", serv.GetProductList)
	http.HandleFunc("/login", serv.Login)
	http.HandleFunc("/register", serv.Register)
	http.HandleFunc("/add-to-cart", serv.AddToCart)
	http.HandleFunc("/delete-from-cart", serv.DeleteFromCart)
	http.HandleFunc("/cart-list", serv.GetCartList)

	http.ListenAndServe(":8090", nil)
}
