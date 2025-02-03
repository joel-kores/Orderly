package categories

import (
	"Orderly/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) CreateCategory(category *models.Category) error {
	return r.DB.Create(category).Error
}

func (r *CategoryRepository) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := r.DB.Preload("Children").First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := r.DB.Preload("Children").Where("parent_id IS NULL").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) UpdateCategory(category *models.Category) error {
	return r.DB.Save(category).Error
}

func (r *CategoryRepository) DeleteCategory(id uint) error {
	return r.DB.Delete(&models.Category{}, id).Error
}
