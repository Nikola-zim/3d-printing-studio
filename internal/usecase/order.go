package usecase

import (
	"context"
	"github.com/Nikola-zim/3d-printing-studio/internal/entity"
)

type OrderManager struct {
}

func NewOrderManager() *OrderManager {
	return &OrderManager{}
}

func (m OrderManager) GetOrders(ctx context.Context, ID int64) ([]entity.Order, error) {
	orders := []entity.Order{
		{
			Info: "smth",
		},
		{
			Info: "smth2",
		},
	}
	return orders, nil
}
