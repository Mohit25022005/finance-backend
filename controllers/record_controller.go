package controllers

import (
	"net/http"
	"strconv"

	"finance-backend/models"
	"finance-backend/services"

	"github.com/gin-gonic/gin"
)

// CreateRecord godoc
// @Summary Create record
// @Description Create a financial record (Admin only)
// @Tags Records
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param record body models.CreateRecordRequest true "Record Data"
// @Success 201 {object} models.RecordResponse
// @Failure 400 {object} map[string]string
// @Router /records [post]
func CreateRecord(svc services.RecordService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.CreateRecordRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		userIDVal, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		userID := userIDVal.(uint)

		record, err := svc.Create(userID, req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, record.ToResponse())
	}
}

// GetRecords godoc
// @Summary Get records
// @Description Get records with filtering and pagination
// @Tags Records
// @Security BearerAuth
// @Produce json
// @Param type query string false "Record type"
// @Param category query string false "Category"
// @Param search query string false "Search"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /records [get]
func GetRecords(svc services.RecordService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.RecordFilterParams

		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query params"})
			return
		}

		records, total, err := svc.GetFiltered(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch records"})
			return
		}

		responses := make([]models.RecordResponse, len(records))
		for i, r := range records {
			responses[i] = r.ToResponse()
		}

		c.JSON(http.StatusOK, gin.H{
			"total":   total,
			"page":    params.Page,
			"limit":   params.Limit,
			"records": responses,
		})
	}
}

// UpdateRecord godoc
// @Summary Update record
// @Description Update a financial record (Admin only)
// @Tags Records
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Record ID"
// @Param record body models.UpdateRecordRequest true "Updated Data"
// @Success 200 {object} models.RecordResponse
// @Failure 400 {object} map[string]string
// @Router /records/{id} [put]
func UpdateRecord(svc services.RecordService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		var req models.UpdateRecordRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		userID := c.GetUint("user_id")

		record, err := svc.Update(uint(id), userID, req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, record.ToResponse())
	}
}

// DeleteRecord godoc
// @Summary Delete record
// @Description Delete a record (Admin only)
// @Tags Records
// @Security BearerAuth
// @Produce json
// @Param id path int true "Record ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /records/{id} [delete]
func DeleteRecord(svc services.RecordService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		userID := c.GetUint("user_id")

		if err := svc.Delete(uint(id), userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "record deleted"})
	}
}