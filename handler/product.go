package handler

import (
	// golang package
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

// GetProductList is a function to handle get product list request
// it accepts http.ResponseWriter and pointer of http.Request as parameters
// it returns status code 200 when success
// otherwise it returns status code 500 when error occurs
func (handler *Handler) GetProductList(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctxHandler, _ := context.WithTimeout(context.Background(), time.Duration(time.Minute*1))

	productCategoryIDStr := req.URL.Query().Get("product_category_id")
	var productCategoryID int64
	if productCategoryIDStr != "" {
		productCategoryID, _ = strconv.ParseInt(productCategoryIDStr, 10, 64)
	}

	result, err := handler.Usecase.GetProductList(ctxHandler, productCategoryID)
	if err != nil {
		log.Println("Handler GetProductList handler.Usecase.GetProductList err: ", err)
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
