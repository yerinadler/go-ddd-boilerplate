package application

import (
	"context"

	"github.com/yerinadler/go-ddd/internal/domain/order"
)

type OrderApplicationService struct {
	repository order.OrderRepository
}

func NewOrderApplicationService(repo order.OrderRepository) *OrderApplicationService {
	return &OrderApplicationService{
		repository: repo,
	}
}

func (svc *OrderApplicationService) GetOrderById(ctx context.Context, id string) (*OrderDTO, error) {
	order, err := svc.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &OrderDTO{
		Id:          order.Id,
		CustomerId:  order.CustomerId,
		Status:      order.Status,
		TotalAmount: order.CalculateTotalAmount(),
		CreatedAt:   order.CreatedAt,
	}, nil
}

func (svc *OrderApplicationService) CreateOrder(ctx context.Context, customer_id string) error {
	order := order.NewOrder(customer_id)
	err := svc.repository.Save(ctx, order)
	if err != nil {
		return err
	}
	return nil
}

func (svc *OrderApplicationService) MarkOrderAsPaid(ctx context.Context, order_id string) error {
	order, err := svc.repository.GetById(ctx, order_id)
	if err != nil {
		return err
	}

	err = order.MarkAsPaid()

	if err != nil {
		return err
	}

	err = svc.repository.Save(ctx, order)

	if err != nil {
		return err
	}

	return nil
}
