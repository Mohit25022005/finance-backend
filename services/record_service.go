package services

import (
	"errors"
	"finance-backend/models"
	"finance-backend/repository"
)

func CreateRecord(record *models.Record) error {
	if record.Amount <= 0 {
		return errors.New("amount must be positive")
	}

	if record.Type != "income" && record.Type != "expense" {
		return errors.New("invalid type")
	}

	return repository.CreateRecord(record)
}

func GetRecords() ([]models.Record, error) {
	return repository.GetAllRecords()
}