package main

import (
	"encoding/json"
	"log"
	"payment/queue"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type Order struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ProductId string `json:"product_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}

func notifyPayment(order Order, ch *amqp.Channel) {
	json, _ := json.Marshal(order)

	queue.Notify(json, "payment_ex", "", ch)
	log.Println(json)
}

func main() {
	in := make(chan []byte)

	connection := queue.Connect()
	queue.StartConsuming("order_queue", connection, in)

	log.Println("IN", in)

	var order Order
	for payload := range in {
		json.Unmarshal(payload, &order)

		order.Status = "approved"
		notifyPayment(order, connection)

		log.Println(string(payload))
	}
}
