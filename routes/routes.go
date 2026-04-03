package routes

import (
	"finance-backend/controllers"
	"finance-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")
	
	api.POST("/login", controllers.Login)
	api.Use(middleware.AuthMiddleware())

	api.POST("/records",
		middleware.RoleMiddleware("admin"),
		controllers.CreateRecord)

	api.GET("/records",
		middleware.RoleMiddleware("admin", "analyst", "viewer"),
		controllers.GetRecords)

	api.GET("/dashboard",
		middleware.RoleMiddleware("admin", "analyst"),
		controllers.GetDashboard)

	api.POST("/users",
		middleware.RoleMiddleware("admin"),
		controllers.CreateUser)

	api.GET("/users",
		middleware.RoleMiddleware("admin"),
		controllers.GetUsers)

	api.PUT("/users/:id",
		middleware.RoleMiddleware("admin"),
		controllers.UpdateUser)

	api.DELETE("/users/:id",
		middleware.RoleMiddleware("admin"),
		controllers.DeleteUser)

	api.PUT("/records/:id",
		middleware.RoleMiddleware("admin"),
		controllers.UpdateRecord)

	api.DELETE("/records/:id",
		middleware.RoleMiddleware("admin"),
		controllers.DeleteRecord)
}
