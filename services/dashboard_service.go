package services

import (
	"finance-backend/config"
	"finance-backend/models"
)

func GetSummary() (map[string]float64, error) {
	var income float64
	var expense float64

	config.DB.Model(&models.Record{}).
		Where("type = ?", "income").
		Select("SUM(amount)").Scan(&income)

	config.DB.Model(&models.Record{}).
		Where("type = ?", "expense").
		Select("SUM(amount)").Scan(&expense)

	return map[string]float64{
		"income":  income,
		"expense": expense,
		"net":     income - expense,
	}, nil
}