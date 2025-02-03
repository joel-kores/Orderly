package order_items_services

import (
	"Orderly/internal/models"
	orderitem "Orderly/internal/repositories/order_items"
	"errors"
)

type OrderItemService struct {
	repo *orderitem.OrderItemRepository
}

func NewOrderItemService(repo *orderitem.OrderItemRepository) *OrderItemService {
	return &OrderItemService{repo: repo}
}

func (s *OrderItemService) CreateOrderItem(orderItem *models.OrderItem) error {
	if orderItem.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	return s.repo.Create(orderItem)
}

func (s *OrderItemService) GetAllOrderItems() ([]models.OrderItem, error) {
	return s.repo.GetAll()
}

func (s *OrderItemService) GetOrderItemByID(id uint) (*models.OrderItem, error) {
	return s.repo.GetByID(id)
}

func (s *OrderItemService) UpdateOrderItem(orderItem *models.OrderItem) error {
	return s.repo.Update(orderItem)
}

func (s *OrderItemService) DeleteOrderItem(id uint) error {
	return s.repo.Delete(id)
}
