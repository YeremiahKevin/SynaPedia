package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func (handler *Handler) GetProductList(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	productCategoryIDStr := req.URL.Query().Get("product_category_id")
	var productCategoryID int64
	if productCategoryIDStr != "" {
		productCategoryID, _ = strconv.ParseInt(productCategoryIDStr, 10, 64)
	}

	ctxHandler, _ := context.WithTimeout(context.Background(), time.Duration(time.Minute*1))

	result, err := handler.Usecase.GetProductList(ctxHandler, productCategoryID)
	if err != nil {
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
