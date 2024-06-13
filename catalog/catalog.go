package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Products struct {
	Products []Product
}

var productsUrl string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	productsUrl = os.Getenv("PRODUCT_URL")
}

func loadProducts() ([]Product, error) {
	response, err := http.Get(productsUrl + "/products")
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var products Products
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return products.Products, nil
}

func listProducts(w http.ResponseWriter, r *http.Request) {
	products, err := loadProducts()
	if err != nil {
		http.Error(w, "Error loading products", http.StatusInternalServerError)
		fmt.Println("Error loading products:", err)
		return
	}

	t, err := template.ParseFiles("templates/catalog.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		fmt.Println("Error parsing template:", err)
		return
	}

	if err := t.Execute(w, products); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		fmt.Println("Error executing template:", err)
	}
}

func showProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := http.Get(productsUrl + "/products/" + vars["id"])
	if err != nil {
		http.Error(w, "Error fetching product", http.StatusInternalServerError)
		fmt.Println("Error fetching product:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		fmt.Println("Error reading response body:", err)
		return
	}

	var product Product
	if err := json.Unmarshal(data, &product); err != nil {
		http.Error(w, "Error unmarshalling product data", http.StatusInternalServerError)
		fmt.Println("Error unmarshalling product data:", err)
		return
	}

	t, err := template.ParseFiles("templates/view.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		fmt.Println("Error parsing template:", err)
		return
	}

	if err := t.Execute(w, product); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		fmt.Println("Error executing template:", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", listProducts)
	r.HandleFunc("/products/{id}", showProducts)
	fmt.Println("Server running at http://localhost:8082")

	if err := http.ListenAndServe(":8082", r); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
