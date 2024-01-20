package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
)

func main() {
	InitializeStorage()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to little Heureka API. For openAPI documentation go to /docs"))
	})

	r.Get("/category", getCategory)
	r.Get("/categories", getCategories)

	r.Get("/product", getProduct)
	r.Get("/products", getProducts)

	r.Get("/offer", getOffer)
	r.Get("/offers", getOffers)

	r.Get("/docs", serveOpenAPI)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)

}
