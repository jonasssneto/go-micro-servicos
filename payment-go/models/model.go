package model

type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float32 `json:"price,string"`
}

type Order struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ProductId string `json:"productId"`
	Status    string `json:"status"`
	CreatedAt string `json:"created"`
}

type OrderPayload struct {
	ProductId   int    `json:"productId"`
	OrderDate   string `json:"order_date"`
	OrderStatus string `json:"order_status"`
}

type OrderResponse struct {
	Id        int    `json:"id"`
	ProductId int    `json:"productId"`
	OrderDate string `json:"order_date"`
}
