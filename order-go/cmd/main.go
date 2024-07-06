package main

import (
	"encoding/json"
	"log"
	"order/config"
	model "order/models"
	"order/repository"
	"order/services/db"
	"order/services/queue"
	"time"
)

func main() {
	config.LoadEnv()

	in := make(chan []byte)
	connection := queue.Connect()
	postgres, err := db.Connect()

	log.Println("[!] Running order service...")

	if err != nil {
		log.Fatal(err)
	}

	OrderRepository := repository.NewOrderRepository(postgres)

	queue.StartConsuming("checkout_queue", connection, in)
	for payload := range in {
		log.Println("Order:", string(payload))

		var orderPayload model.OrderPayload
		err := json.Unmarshal(payload, &orderPayload)
		if err != nil {
			panic(err)
		}

		orderPayload.OrderDate = time.Now().Format("01-02-2006")
		orderPayload.OrderStatus = "pending"

		data, err := OrderRepository.CreateOrder(orderPayload.ProductId, orderPayload.OrderDate)

		if err != nil {
			log.Println("Error creating order:", err)
			panic(err)
		}

		orderResponse := model.OrderResponse{
			Id:        data.Id,
			ProductId: data.ProductId,
			OrderDate: data.CreatedAt,
		}

		OrderRepository.PublishOrder(orderResponse, connection)
	}
}
