package repository

import (
	"finance-backend/config"
	"finance-backend/models"
)

func CreateRecord(record *models.Record) error {
	return config.DB.Create(record).Error
}

func GetAllRecords() ([]models.Record, error) {
	var records []models.Record
	err := config.DB.Find(&records).Error
	return records, err
}