package main

import (
	// golang package
	"log"
	"net/http"
	"os"
	"strings"

	// external package
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kataras/jwt"

	// internal package
	"SynaPedia/handler"
	"SynaPedia/repository"
	"SynaPedia/usecase"
)

func main() {
	log.Println("Running handler...")

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	repo := repository.NewRepository()
	uc := usecase.NewUsecase(repo)
	serv := handler.NewHandler(uc)

	http.Handle("/product-list", http.HandlerFunc(serv.GetProductList))
	http.Handle("/login", http.HandlerFunc(serv.Login))
	http.Handle("/register", http.HandlerFunc(serv.Register))
	http.Handle("/add-to-cart", middleware(http.HandlerFunc(serv.AddToCart)))
	http.Handle("/delete-from-cart", middleware(http.HandlerFunc(serv.DeleteFromCart)))
	http.Handle("/cart-list", middleware(http.HandlerFunc(serv.GetCartList)))
	http.Handle("/create-order", middleware(http.HandlerFunc(serv.CreateOrder)))

	http.ListenAndServe(":8080", nil)
}

// middleware is a function to intercept HTTP request
// get authorization token from header and check validity
// it returns function http.Handler when token valid
// otherwise it returns error unauthorized
func middleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		reqToken := req.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) == 2 {
			reqToken = splitToken[1]
		}

		_, err := jwt.Verify(jwt.HS256, []byte(os.Getenv("JWT_KEY")), []byte(reqToken))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		nextHandler.ServeHTTP(w, req)
	})
}
