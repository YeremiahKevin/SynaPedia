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

	http.ListenAndServe(":8090", nil)
}
