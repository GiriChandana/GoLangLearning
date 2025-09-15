package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: homepage")
	fmt.Fprintf(w, "welcome to homepage")
}
func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: returnAllProducts")
	json.NewEncoder(w).Encode(Products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	key := r.URL.Path[len("/product/"):]
	for _, product := range Products {
		if string(product.Id) == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}
func handleRequests() {
	http.HandleFunc("/products", returnAllProducts)
	http.HandleFunc("/product/", getProduct)
	http.HandleFunc("/", homePage)
	http.ListenAndServe("localhost:8080", nil)
}
func main() {
	Products = []Product{
		{Id: "1", Name: "Chair", Quantity: 100, Price: 100.00},
		{Id: "2", Name: "Desk", Quantity: 200, Price: 200.00},
	}
	handleRequests()
}
