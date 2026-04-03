package controllers

import (
	"errors"
	"net/http"

	"finance-backend/config"
	"finance-backend/models"
	"finance-backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LoginRequest represents login input
type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// Login godoc
// @Summary Login user
// @Description Login using email and get JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body controllers.LoginRequest true "Login Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /login [post]
func Login(c *gin.Context) {
	var req LoginRequest

	// ✅ Validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// ✅ Fetch user from DB
	var user models.User
	err := config.DB.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid credentials",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "database error",
		})
		return
	}

	// ✅ Check if user is active
	if !user.IsActive {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "user account is inactive",
		})
		return
	}

	// ✅ Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Role.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	// ✅ Success response
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}