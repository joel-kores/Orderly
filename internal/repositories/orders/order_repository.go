package repository

import (
	"Orderly/internal/models"
	"context"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, order *models.Order) error
	GetByID(ctx context.Context, id uint) (*models.Order, error)
	GetAll(ctx context.Context) ([]models.Order, error)
	Update(ctx context.Context, order *models.Order) error
	Delete(ctx context.Context, id uint) error
}

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepositoryImpl{db: db}
}

func (r *orderRepositoryImpl) Create(ctx context.Context, order *models.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *orderRepositoryImpl) GetByID(ctx context.Context, id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.WithContext(ctx).Preload("Items").First(&order, id).Error
	return &order, err
}

func (r *orderRepositoryImpl) GetAll(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.WithContext(ctx).Preload("Items").Find(&orders).Error
	return orders, err
}

func (r *orderRepositoryImpl) Update(ctx context.Context, order *models.Order) error {
	return r.db.WithContext(ctx).Save(order).Error
}

func (r *orderRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Order{}, id).Error
}
