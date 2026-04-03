package controllers

import (
	"net/http"

	"finance-backend/services"

	"github.com/gin-gonic/gin"
)

// GetDashboard godoc
// @Summary Get dashboard data
// @Description Get financial summary and analytics
// @Tags Dashboard
// @Security BearerAuth
// @Produce json
// @Success 200 {object} services.DashboardData
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /dashboard [get]
func GetDashboard(svc services.DashboardService) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Extract user_id
		userIDVal, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		userID, ok := userIDVal.(uint)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid user context",
			})
			return
		}

		// Call service
		data, err := svc.GetDashboardData(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to fetch dashboard data",
			})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}