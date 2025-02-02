package order_item

import (
	"Orderly/internal/models"
	"gorm.io/gorm"
)

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}

func (r *OrderItemRepository) Create(orderItem *models.OrderItem) error {
	return r.db.Create(orderItem).Error
}

func (r *OrderItemRepository) GetAll() ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	err := r.db.Preload("Product").Preload("Order").Find(&orderItems).Error
	return orderItems, err
}

func (r *OrderItemRepository) GetByID(id uint) (*models.OrderItem, error) {
	var orderItem models.OrderItem
	err := r.db.Preload("Product").Preload("Order").First(&orderItem, id).Error
	if err != nil {
		return nil, err
	}
	return &orderItem, nil
}

func (r *OrderItemRepository) Update(orderItem *models.OrderItem) error {
	return r.db.Save(orderItem).Error
}

func (r *OrderItemRepository) Delete(id uint) error {
	return r.db.Delete(&models.OrderItem{}, id).Error
}
