package products

import (
	"checkout/model"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetProductById(productId int) (*model.Product, error) {
	productsUrl := os.Getenv("PRODUCTS_API_URL")

	url := productsUrl + "/products/" + strconv.Itoa(productId)

	response, err := http.Get(url)

	if err != nil {
		log.Println("Error fetching product:", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errorResponse := struct {
			Error string `json:"error"`
		}{
			Error: "product not found",
		}

		errorMsg, _ := json.Marshal(errorResponse)
		log.Println("Error fetching product:", response.StatusCode)
		return nil, errors.New(string(errorMsg))
	}

	var product model.Product
	if err := json.NewDecoder(response.Body).Decode(&product); err != nil {
		log.Println("Error decoding product response:", err)
		return nil, err
	}

	return &product, nil
}
