package repository

import (
	"Orderly/internal/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) GetAll() (
	[]models.GetProducts,
	error,
) {
	var products []models.Product
	err := r.db.Preload("Category").Find(&products).Error
	if err != nil {
		return nil, err
	}

	// Transform []Product → []GetProducts
	var response []models.GetProducts
	for _, p := range products {
		response = append(response, models.GetProducts{
			CommonFields: p.CommonFields,
			Name:         p.Name,
			Price:        p.Price,
			CategoryID:   p.CategoryID,
			Category:     p.Category, // Make sure Category is loaded if needed
		})
	}

	return response, nil
}

func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
