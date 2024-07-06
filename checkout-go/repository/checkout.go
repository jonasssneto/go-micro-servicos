package repository

import (
	"checkout/model"
	"checkout/services/products"
	"checkout/services/queue"
	"encoding/json"
	"errors"
	"log"

	"github.com/streadway/amqp"
)

type CheckoutRepository struct {
	queue *amqp.Channel
}

func NewCheckoutRepository(queue *amqp.Channel) CheckoutRepository {
	return CheckoutRepository{
		queue: queue,
	}
}

func (cr *CheckoutRepository) Checkout(order model.Order) ([]byte, error) {
	data, err := json.Marshal(order)

	if err != nil {
		return nil, err
	}

	product, err := products.GetProductById(order.ProductId)

	if err != nil {
		return nil, errors.New("product not found")
	}

	log.Println("Product: ", product)

	queue.Publish(data, "checkout_ex", "", cr.queue)

	return data, nil
}
