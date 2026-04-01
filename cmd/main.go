package main

import (
	"finance-backend/config"
	"finance-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8084")
}