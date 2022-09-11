package order

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	Created   OrderStatus = "CREATED"
	Paid      OrderStatus = "PAID"
	Cancelled OrderStatus = "CANCELLED"
)

type Order struct {
	Id         string      `json:"id" bson:"id"`
	CustomerId string      `json:"customerId" bson:"customerId"`
	OrderItems []OrderItem `json:"orderItems" bson:"orderItems"`
	Status     OrderStatus `json:"status" bson:"status"`
	CreatedAt  time.Time   `json:"createdAt" bson:"createdAt"`
}

func NewOrder(customer_id string) *Order {
	order := Order{
		Id:         uuid.NewString(),
		CustomerId: customer_id,
		OrderItems: []OrderItem{},
		Status:     Created,
		CreatedAt:  time.Now(),
	}
	return &order
}

func (o *Order) AddOrderItem(product_id string, quantity int, price int) {
	order_item := OrderItem{
		ProductId: product_id,
		Quantity:  quantity,
		Price:     price,
	}
	o.OrderItems = append(o.OrderItems, order_item)
}

func (o *Order) MarkAsPaid() error {
	if o.Status != Created {
		return errors.New("can not mark the order as paid : invalid status")
	}
	o.Status = Paid
	return nil
}

func (o *Order) MarkAsCancelled() error {
	o.Status = Cancelled
	return nil
}

func (o *Order) CalculateTotalAmount() int {
	if len(o.OrderItems) == 0 {
		return 0
	}

	total := 0

	for _, item := range o.OrderItems {
		total += item.Price
	}

	return total
}
