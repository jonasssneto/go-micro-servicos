package main

import (
	"catalog/db"
	"catalog/queue"
	"encoding/json"
	"flag"
	"log"
	"time"

	"github.com/joho/godotenv"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/streadway/amqp"
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

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}

func createOrder(payload []byte) Order {
	var order Order
	json.Unmarshal(payload, &order)

	uuid, _ := uuid.NewV4()
	order.Uuid = uuid.String()
	order.Status = "pending"
	order.CreatedAt = time.Now().String()
	saveOrder(order)

	return order
}

func saveOrder(order Order) {
	json, _ := json.Marshal(order)

	connection := db.Connect()

	err := connection.Set(order.Uuid, json, 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Order saved with success")
}

func notifyOrderCreated(order Order, ch *amqp.Channel) {
	json, _ := json.Marshal(order)

	queue.Notify(json, "order_ex", "", ch)
}

func main() {
	var param string
	flag.StringVar(&param, "opt", "", "Usage")
	flag.Parse()

	in := make(chan []byte)
	connection := queue.Connect()

	switch param {
	case "checkout":
		queue.StartConsuming("checkout_queue", connection, in)
		for payload := range in {
			notifyOrderCreated(createOrder(payload), connection)
			log.Println(string(payload))
		}
	case "payment":
		queue.StartConsuming("payment_queue", connection, in)
		var order Order
		for payload := range in {
			json.Unmarshal(payload, &order)
			saveOrder(order)
			log.Println("Payment: ", string(payload))
		}
	}
}
