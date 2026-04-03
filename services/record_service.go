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

func UpdateRecord(id uint, updated *models.Record) error {
	record, err := repository.GetRecordByID(id)
	if err != nil {
		return errors.New("record not found")
	}

	// validation
	if updated.Amount <= 0 {
		return errors.New("amount must be positive")
	}

	if updated.Type != "income" && updated.Type != "expense" {
		return errors.New("invalid type")
	}

	// update fields
	record.Amount = updated.Amount
	record.Type = updated.Type
	record.Category = updated.Category
	record.Notes = updated.Notes

	return repository.UpdateRecord(&record)
}

func DeleteRecord(id uint) error {
	_, err := repository.GetRecordByID(id)
	if err != nil {
		return errors.New("record not found")
	}

	return repository.DeleteRecord(id)
}