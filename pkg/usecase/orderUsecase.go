package usecase

import (
	"github.com/akshayur04/project-ecommerce/pkg/domain"
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
)

type OrderUseCase struct {
	orderRepo interfaces.OrderRepository
}

func NewOrderUseCase(orderRepo interfaces.OrderRepository) services.OrderUseCase {
	return &OrderUseCase{
		orderRepo: orderRepo,
	}
}

func (c *OrderUseCase) OrderAll(id, paymentTypeId int) (domain.Orders, error) {
	order, err := c.orderRepo.OrderAll(id, paymentTypeId)
	return order, err
}

func (c *OrderUseCase) UserCancelOrder(orderId, userId int) error {
	err := c.orderRepo.UserCancelOrder(orderId, userId)
	return err
}
