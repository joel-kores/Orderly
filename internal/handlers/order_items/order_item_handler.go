package order_items

import (
	"Orderly/internal/models"
	"Orderly/internal/services/order_items"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderItemHandler struct {
	service *order_items_services.OrderItemService
}

func NewOrderItemHandler(service *order_items_services.OrderItemService) *OrderItemHandler {
	return &OrderItemHandler{service: service}
}

func (h *OrderItemHandler) CreateOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	if err := c.ShouldBindJSON(&orderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateOrderItem(&orderItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, orderItem)
}

func (h *OrderItemHandler) GetAllOrderItems(c *gin.Context) {
	orderItems, err := h.service.GetAllOrderItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderItems)
}

func (h *OrderItemHandler) GetOrderItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	orderItem, err := h.service.GetOrderItemByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
		return
	}
	c.JSON(http.StatusOK, orderItem)
}

func (h *OrderItemHandler) UpdateOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	if err := c.ShouldBindJSON(&orderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateOrderItem(&orderItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderItem)
}

func (h *OrderItemHandler) DeleteOrderItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.service.DeleteOrderItem(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order item deleted"})
}
