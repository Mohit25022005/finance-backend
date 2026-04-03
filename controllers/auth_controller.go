package controllers

import (
	"finance-backend/config"
	"finance-backend/models"
	"finance-backend/utils"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email string `json:"email"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	var user models.User
	err := config.DB.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}