package application

import (
	"time"

	"github.com/yerinadler/go-ddd/internal/domain/order"
)

type CreateOrderRequestDTO struct {
	CustomerId string `json:"customerId"`
}

type OrderDTO struct {
	Id          string            `json:"id"`
	CustomerId  string            `json:"customerId"`
	TotalAmount int               `json:"totalAmount"`
	Status      order.OrderStatus `json:"orderStatus"`
	CreatedAt   time.Time         `json:"createdAt"`
}
