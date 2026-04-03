package controllers

import (
	"finance-backend/models"
	"finance-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRecord(c *gin.Context) {
	var record models.Record

	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateRecord(&record)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, record)
}

func GetRecords(c *gin.Context) {

	typeFilter := c.Query("type")
	categoryFilter := c.Query("category")

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	// safety checks
	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	filters := map[string]string{
		"type":     typeFilter,
		"category": categoryFilter,
	}

	records, err := services.GetRecords(filters, page, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"page":    page,
		"limit":   limit,
		"records": records,
	})
}

func UpdateRecord(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var record models.Record
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := services.UpdateRecord(uint(id), &record)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "record updated"})
}

func DeleteRecord(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := services.DeleteRecord(uint(id))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "record deleted"})
}