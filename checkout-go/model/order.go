package model

type Order struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ProductId int    `json:"productId"`
}

type CheckoutResponse struct {
	Data    []byte
	Product *Product
}
