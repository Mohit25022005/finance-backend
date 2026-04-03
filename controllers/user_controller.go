package controllers

import (
	"net/http"
	"strconv"

	"finance-backend/models"
	"finance-backend/services"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary Create user
// @Description Create a new user (Admin only)
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User Data"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} map[string]string
// @Router /users [post]
func CreateUser(svc services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.CreateUserRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		user, err := svc.Create(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, user.ToResponse())
	}
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get a user by ID (Admin only)
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.UserResponse
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func GetUserByID(svc services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}

		user, err := svc.GetByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user.ToResponse())
	}
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users (Admin only)
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.UserResponse
// @Failure 500 {object} map[string]string
// @Router /users [get]
func GetUsers(svc services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := svc.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
			return
		}

		res := make([]models.UserResponse, len(users))
		for i, u := range users {
			res[i] = u.ToResponse()
		}

		c.JSON(http.StatusOK, res)
	}
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user details (Admin only)
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UpdateUserRequest true "Updated Data"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} map[string]string
// @Router /users/{id} [put]
func UpdateUser(svc services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}

		var req models.UpdateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		user, err := svc.Update(uint(id), req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user.ToResponse())
	}
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user (Admin only)
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /users/{id} [delete]
func DeleteUser(svc services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}

		if err := svc.Delete(uint(id)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
	}
}