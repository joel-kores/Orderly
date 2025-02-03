package services

import (
	"Orderly/internal/models"
	repository "Orderly/internal/repositories/orders"
	"context"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *models.Order) error
	GetOrderByID(ctx context.Context, id uint) (*models.Order, error)
	GetAllOrders(ctx context.Context) ([]models.Order, error)
	UpdateOrder(ctx context.Context, order *models.Order) error
	DeleteOrder(ctx context.Context, id uint) error
}

type orderServiceImpl struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderServiceImpl{repo: repo}
}

func (s *orderServiceImpl) CreateOrder(ctx context.Context, order *models.Order) error {
	return s.repo.Create(ctx, order)
}

func (s *orderServiceImpl) GetOrderByID(ctx context.Context, id uint) (*models.Order, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *orderServiceImpl) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	return s.repo.GetAll(ctx)
}

func (s *orderServiceImpl) UpdateOrder(ctx context.Context, order *models.Order) error {
	return s.repo.Update(ctx, order)
}

func (s *orderServiceImpl) DeleteOrder(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
