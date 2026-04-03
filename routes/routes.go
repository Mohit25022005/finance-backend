package routes

import (
	"os"

	"finance-backend/config"
	"finance-backend/controllers"
	"finance-backend/middleware"
	"finance-backend/repository"
	"finance-backend/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine) {

	// Global middleware
	r.Use(middleware.RateLimitMiddleware())

	// Swagger only in non-production
	if os.Getenv("APP_ENV") != "production" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// --- Initialize dependencies ---
	db := config.DB

	userRepo := repository.NewUserRepository(db)
	recordRepo := repository.NewRecordRepository(db)

	userService := services.NewUserService(userRepo)
	recordService := services.NewRecordService(recordRepo)
	dashboardService := services.NewDashboardService(db)

	// --- Routes ---
	api := r.Group("/api/v1")

	// Public
	api.POST("/login", controllers.Login)

	// Protected
	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleware())

	// ----- RECORDS -----
	auth.POST("/records",
		middleware.RoleMiddleware("admin"),
		controllers.CreateRecord(recordService),
	)

	auth.GET("/records",
		middleware.RoleMiddleware("admin", "analyst", "viewer"),
		controllers.GetRecords(recordService),
	)

	auth.PUT("/records/:id",
		middleware.RoleMiddleware("admin"),
		controllers.UpdateRecord(recordService),
	)

	auth.DELETE("/records/:id",
		middleware.RoleMiddleware("admin"),
		controllers.DeleteRecord(recordService),
	)

	// ----- DASHBOARD -----
	auth.GET("/dashboard",
		middleware.RoleMiddleware("admin", "analyst"),
		controllers.GetDashboard(dashboardService),
	)

	// ----- USERS -----
	auth.POST("/users",
		middleware.RoleMiddleware("admin"),
		controllers.CreateUser(userService),
	)

	auth.GET("/users",
		middleware.RoleMiddleware("admin"),
		controllers.GetUsers(userService),
	)

	auth.GET("/users/:id",
		middleware.RoleMiddleware("admin"),
		controllers.GetUserByID(userService),
	)

	auth.PUT("/users/:id",
		middleware.RoleMiddleware("admin"),
		controllers.UpdateUser(userService),
	)

	auth.DELETE("/users/:id",
		middleware.RoleMiddleware("admin"),
		controllers.DeleteUser(userService),
	)

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}