package Handlers

import (
	"Orderly/internal/models"
	"Orderly/internal/services/categories"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryHandler struct {
	CategoryService *Services.CategoryService
}

func NewCategoryHandler(categoryService *Services.CategoryService) *CategoryHandler {
	return &CategoryHandler{CategoryService: categoryService}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.CategoryService.CreateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var categoryID uint
	if _, err := fmt.Sscanf(id, "%d", &categoryID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}
	category, err := h.CategoryService.GetCategoryByID(categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.CategoryService.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.CategoryService.UpdateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var categoryID uint
	if _, err := fmt.Sscanf(id, "%d", &categoryID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}
	if err := h.CategoryService.DeleteCategory(categoryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
