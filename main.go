package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"strconv"
)

func main() {
	InitializeStorage()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to little Heureka API"))
	})

	r.Get("/categories", GetCategories)
	r.Get("/category", GetCategory)

	r.Get("/products", GetProducts)
	r.Get("/product", GetProduct)

	r.Get("/offer", GetOffer)
	r.Get("/offers", GetOffers)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)

}

func GetCategory(w http.ResponseWriter, r *http.Request) {
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

func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories := ReadCategories()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(categories)
	if err != nil {
		log.Fatal(err)
	}
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
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

func GetProducts(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(products[offset:limit])
}

func GetOffer(w http.ResponseWriter, r *http.Request) {
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

func GetOffers(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(offers[offset:limit])

}
