package main

import (
	"catalog/queue"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float32 `json:"price,string"`
}

type Order struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ProductId string `json:"product_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

var productsUrl string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	productsUrl = os.Getenv("PRODUCT_URL")
}

func main() {
	in := make(chan []byte)

	connection := queue.Connect()
	queue.StartConsuming(connection, in)

	for payload := range in {
		log.Println(string(payload))
	}
}
