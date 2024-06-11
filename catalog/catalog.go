package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Products struct {
	Products []Product
}

var productsUrl = "http://localhost:8081"

func loadProducts() []Product {
	fmt.Println(productsUrl + "/products")
	response, err := http.Get(productsUrl + "/products")
	if err != nil {
		fmt.Println("erro http:" + err.Error())
	}
	data, _ := io.ReadAll(response.Body)

	var products Products

	json.Unmarshal(data, &products)

	fmt.Println(data)

	return products.Products
}

func listProducts(w http.ResponseWriter, r *http.Request) {
	products := loadProducts()
	t := template.Must((template.ParseFiles("templates/catalog.html")))
	t.Execute(w, products)
}

func showProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := http.Get(productsUrl + "/products" + vars["id"])
	if err != nil {
		fmt.Println("erro http:" + err.Error())
	}
	data, _ := io.ReadAll(response.Body)

	var product Product
	json.Unmarshal(data, &product)

	t := template.Must((template.ParseFiles("templates/view.html")))
	t.Execute(w, product)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", listProducts)
	r.HandleFunc("/products/{id}", showProducts)
	http.ListenAndServe(":8082", r)
}
