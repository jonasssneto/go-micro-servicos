package repository

import (
	"database/sql"
)

type PaymentRepository struct {
	connection *sql.DB
}

func NewPaymentRepository(connection *sql.DB) PaymentRepository {
	return PaymentRepository{
		connection,
	}
}

func (pr *PaymentRepository) ProcessPayment(id int) error {
	query := `UPDATE orders SET order_status = 'success' WHERE id = $1`

	row := pr.connection.QueryRow(query, id)

	err := row.Scan()

	if err != nil {
		return err
	}

	return nil
}
