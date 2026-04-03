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

func GetRecordByID(id uint) (models.Record, error) {
	var record models.Record
	err := config.DB.First(&record, id).Error
	return record, err
}

func UpdateRecord(record *models.Record) error {
	return config.DB.Save(record).Error
}

func DeleteRecord(id uint) error {
	return config.DB.Delete(&models.Record{}, id).Error
}

func GetFilteredRecords(filters map[string]string, page int, limit int) ([]models.Record, error) {
	var records []models.Record

	query := config.DB.Model(&models.Record{})

	// filtering
	if filters["type"] != "" {
		query = query.Where("type = ?", filters["type"])
	}
	if filters["category"] != "" {
		query = query.Where("category = ?", filters["category"])
	}

	// pagination
	offset := (page - 1) * limit

	err := query.Limit(limit).Offset(offset).Find(&records).Error
	return records, err
}