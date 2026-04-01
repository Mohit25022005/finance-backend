package routes

import (
	"finance-backend/controllers"
	"finance-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")

	api.POST("/records",
		middleware.RoleMiddleware("admin"),
		controllers.CreateRecord)

	api.GET("/records",
		middleware.RoleMiddleware("admin", "analyst", "viewer"),
		controllers.GetRecords)

	api.GET("/dashboard",
		middleware.RoleMiddleware("admin", "analyst"),
		controllers.GetDashboard)
}