package Services

import (
	"Orderly/internal/models"
	"Orderly/internal/repositories/categories"
)

type CategoryService struct {
	CategoryRepo *categories.CategoryRepository
}

func NewCategoryService(categoryRepo *categories.CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepo: categoryRepo}
}

// Create a new category
func (s *CategoryService) CreateCategory(category *models.Category) error {
	return s.CategoryRepo.CreateCategory(category)
}

// Get a category by ID
func (s *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	return s.CategoryRepo.GetCategoryByID(id)
}

// Get all categories
func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.CategoryRepo.GetAllCategories()
}

// Update a category
func (s *CategoryService) UpdateCategory(category *models.Category) error {
	return s.CategoryRepo.UpdateCategory(category)
}

// Delete a category
func (s *CategoryService) DeleteCategory(id uint) error {
	return s.CategoryRepo.DeleteCategory(id)
}
