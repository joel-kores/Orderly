package handlers

import (
	"github.com/gin-gonic/gin"
)

// Healthcheck godoc
// @Summary Get health status of app
// @Description Get health status of app
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func Healthcheck(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Welcome to Orderly"})
}
