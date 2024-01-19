package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	InitializeStorage()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to little Heureka API. For openAPI documentation go to /docs"))
	})

	r.Get("/category", GetCategory)
	r.Get("/categories", GetCategories)

	r.Get("/product", GetProduct)
	r.Get("/products", GetProducts)

	r.Get("/offer", GetOffer)
	r.Get("/offers", GetOffers)

	r.Get("/docs", serveOpenAPI)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)

}
