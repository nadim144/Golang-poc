package main

import (
	"log"
	"net/http"
	"wipro-toyota-poc/svc/handlers"
	"wipro-toyota-poc/svc/models"
	"wipro-toyota-poc/svc/repository"
)

func main() {

	// Add few sample products
	repository.AddProduct(models.Product{
		Name:      "Milk",
		UnitPrice: 4.00,
	})
	repository.AddProduct(models.Product{
		Name:      "Bread",
		UnitPrice: 5.00,
	})

	// Routing incoming request
	http.HandleFunc("/", handlers.HandleRequest)

	err := http.ListenAndServe(":8080", nil)

	// Error check
	if err != nil {
		log.Println(err)
		return
	}
}
