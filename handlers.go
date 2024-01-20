package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func getCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	category := ReadCategory(id)
	if (category == Category{}) {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(category)
	if err != nil {
		log.Fatal(err)
	}
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	categories := ReadCategories()
	response := constructResponse(categories, 0, 0, len(categories))

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	product := ReadProduct(id)
	if (product == Product{}) {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		log.Fatal(err)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	strCategoryId := r.URL.Query().Get("categoryId")
	var categoryId *int
	if strCategoryId != "" {
		temp, err := strconv.Atoi(strCategoryId)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}
		categoryId = &temp
	}

	products := ReadProducts(categoryId)

	strOffset := r.URL.Query().Get("offset")
	strLimit := r.URL.Query().Get("limit")

	offset, err := strconv.Atoi(strOffset)
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(strLimit)
	if err != nil || limit == 0 || limit > len(products) {
		limit = len(products)
	}

	if offset > len(products) {
		http.Error(w, "Offset is beyond the number of products", http.StatusBadRequest)
		return
	}

	response := constructResponse(products[offset:limit], offset, limit, len(products))

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
}

func getOffer(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	offer := ReadOffer(id)
	if (offer == Offer{}) {
		http.Error(w, "Offer not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(offer)
	if err != nil {
		log.Fatal(err)
	}
}

func getOffers(w http.ResponseWriter, r *http.Request) {
	strProductId := r.URL.Query().Get("productId")
	var productId *int
	if strProductId != "" {
		temp, err := strconv.Atoi(strProductId)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}
		productId = &temp
	}

	offers := ReadOffers(productId)

	strOffset := r.URL.Query().Get("offset")
	strLimit := r.URL.Query().Get("limit")

	offset, err := strconv.Atoi(strOffset)
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(strLimit)
	if err != nil || limit == 0 || limit > len(offers) {
		limit = len(offers)
	}

	if offset > len(offers) {
		http.Error(w, "Offset is beyond the number of offers", http.StatusBadRequest)
		return
	}

	response := constructResponse(offers[offset:limit], offset, limit, len(offers))

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

}

// OpenAPI 3.0.0
func serveOpenAPI(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./openapi.json")
}

func constructResponse(data interface{}, offset int, limit int, total int) Response {
	pageInfo := PageInfo{
		Offset: offset,
		Limit:  limit,
		Total:  total,
	}
	response := Response{
		PageInfo: pageInfo,
		Data:     data,
	}
	return response
}
