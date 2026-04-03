package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"finance-backend/config"
	"finance-backend/models"
	"finance-backend/routes"

	_ "finance-backend/docs" // Swagger docs

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title Finance Backend API
// @version 1.0
// @description Finance dashboard backend API

// @host localhost:8084
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter "Bearer <your_token>"
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := run(); err != nil {
		slog.Error("server exited with error", "error", err)
		os.Exit(1)
	}
}

func run() error {
	// --- DB ---
	if err := config.ConnectDB(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := seedAdminUser(); err != nil {
		return fmt.Errorf("failed to seed admin user: %w", err)
	}

	// --- Gin mode ---
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// --- Router ---
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(requestLoggerMiddleware())
	
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // allow all (ok for assignment)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	
	r.SetTrustedProxies(nil)
	// --- Routes (includes Swagger route already) ---
	routes.RegisterRoutes(r)

	// --- Server config ---
	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// --- Start server ---
	serverErr := make(chan error, 1)
	go func() {
		slog.Info("server starting", "port", port, "env", os.Getenv("APP_ENV"))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	// --- Graceful shutdown ---
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		return fmt.Errorf("server error: %w", err)
	case sig := <-quit:
		slog.Info("shutdown signal received", "signal", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("graceful shutdown failed: %w", err)
	}

	slog.Info("server shutdown complete")
	return nil
}

// --- Seed Admin ---
func seedAdminUser() error {
	admin := models.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Role:     models.RoleAdmin,
		IsActive: true,
	}

	result := config.DB.Where(models.User{Email: admin.Email}).FirstOrCreate(&admin)
	if result.Error != nil {
		return fmt.Errorf("seed admin user: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		slog.Info("admin user already exists", "email", admin.Email)
	} else {
		slog.Info("admin user created", "email", admin.Email)
	}

	return nil
}

// --- Logger Middleware ---
func requestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		slog.Info("request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency", time.Since(start).String(),
			"ip", c.ClientIP(),
		)
	}
}
