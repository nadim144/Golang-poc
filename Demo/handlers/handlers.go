package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"wipro-toyota-poc/svc/models"
	"wipro-toyota-poc/svc/repository"
)

// HandleRequest dispatches based on HTTP verb
func HandleRequest(w http.ResponseWriter, r *http.Request) {

	log.Println("Request received:", r.Method)

	// Method get call based on Verbs
	switch r.Method {
	case http.MethodGet: //GET
		list(w, r)
		break
	case http.MethodPost: //POST
		add(w, r)
		break
	default:
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		break
	}
}

func list(w http.ResponseWriter, r *http.Request) {

	// Get product and convert into Json using Marshal
	products := repository.GetProducts()
	json, _ := json.Marshal(products)

	//Write to header
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200) // 200 status code for success
	w.Write(json)

	// Print log
	log.Println("Successfully Returned Response:", 200)
}

func add(w http.ResponseWriter, r *http.Request) {

	payload, _ := ioutil.ReadAll(r.Body) //ioutil is standard library to body of request.

	var product models.Product
	err := json.Unmarshal(payload, &product) // payload is a Json

	// Check for errors
	if err != nil || product.Name == "" || product.UnitPrice == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}

	// Get newly added product ID
	product.ID = repository.AddProduct(product)

	// Write to the header
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201) // 201 status code for success
	json, _ := json.Marshal(product)
	w.Write(json)

	// Print log
	log.Println("Successfully Returned Response:", 201)
}
