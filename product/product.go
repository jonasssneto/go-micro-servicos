package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

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

func loadData() ([]byte, error) {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return data, nil
}

func listProducts(w http.ResponseWriter, r *http.Request) {
	data, err := loadData()
	if err != nil {
		http.Error(w, "Error loading products", http.StatusInternalServerError)
		fmt.Println("Error loading products:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func getProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data, err := loadData()
	if err != nil {
		http.Error(w, "Error loading products", http.StatusInternalServerError)
		fmt.Println("Error loading products:", err)
		return
	}

	var products Products
	if err := json.Unmarshal(data, &products); err != nil {
		http.Error(w, "Error unmarshalling products", http.StatusInternalServerError)
		fmt.Println("Error unmarshalling products:", err)
		return
	}

	for _, v := range products.Products {
		if v.Uuid == vars["id"] {
			product, err := json.Marshal(v)
			if err != nil {
				http.Error(w, "Error marshalling product", http.StatusInternalServerError)
				fmt.Println("Error marshalling product:", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(product)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", listProducts).Methods("GET")
	r.HandleFunc("/products/{id}", getProductById).Methods("GET")

	fmt.Println("Server running at http://localhost:8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
