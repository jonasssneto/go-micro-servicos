package repository

import (
	"database/sql"
	"encoding/json"
	"log"
	model "order/models"
	"order/services/queue"

	"github.com/streadway/amqp"
)

type OrderRepository struct {
	connection *sql.DB
}

func NewOrderRepository(connection *sql.DB) OrderRepository {
	return OrderRepository{
		connection,
	}
}

func (pr *OrderRepository) PublishOrder(order model.OrderResponse, ch *amqp.Channel) {
	jsonOrder, _ := json.Marshal(order)

	queue.Publish(jsonOrder, "order_ex", "", ch)
}

func (pr *OrderRepository) CreateOrder(productId int, orderDate string) (*model.Order, error) {
	query := `INSERT INTO orders (product_id, order_date) VALUES ($1, $2) RETURNING id, product_id, order_date`

	row := pr.connection.QueryRow(query, productId, orderDate)

	order := &model.Order{}
	err := row.Scan(&order.Id, &order.ProductId, &order.CreatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return order, nil
}
