package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Get("/", listProducts)
		r.Post("/", createProduct)
		r.Get("/{id}", getProduct)
		r.Put("/{id}", updateProduct)
		r.Delete("/{id}", deleteProduct)
	})

	fmt.Println("Server is listening on port 8088")
	http.ListenAndServe(":8088", r)
}
