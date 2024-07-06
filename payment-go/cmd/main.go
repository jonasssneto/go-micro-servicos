package main

import (
	"encoding/json"
	"log"
	"payment/config"
	model "payment/models"
	"payment/repository"
	"payment/services/db"
	"payment/services/queue"
)

func main() {
	config.LoadEnv()

	in := make(chan []byte)
	connection := queue.Connect()
	postgres, err := db.Connect()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("[!] Running payment service...")

	PaymentRepository := repository.NewPaymentRepository(postgres)

	queue.StartConsuming("order_queue", connection, in)
	for payload := range in {
		log.Println("Processing payment:", string(payload))

		var orderPayload model.OrderResponse
		err := json.Unmarshal(payload, &orderPayload)
		if err != nil {
			log.Println("Error decoding payload:", err)
			continue
		}

		err = PaymentRepository.ProcessPayment(orderPayload.Id)

		log.Println("Payment processed")

		if err != nil {
			log.Println("Error processing payment:", err)
		}
	}
}
