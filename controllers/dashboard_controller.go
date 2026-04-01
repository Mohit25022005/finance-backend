package controllers

import (
	"finance-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDashboard(c *gin.Context) {
	data, _ := services.GetSummary()
	c.JSON(http.StatusOK, data)
}