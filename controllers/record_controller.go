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
	records, _ := services.GetRecords()
	c.JSON(200, records)
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