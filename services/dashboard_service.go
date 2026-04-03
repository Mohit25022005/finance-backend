package services

import (
	"finance-backend/config"
	"finance-backend/models"
	
)

func GetDashboardData() (map[string]interface{}, error) {

	var income float64
	var expense float64

	// totals
	config.DB.Model(&models.Record{}).
		Where("type = ?", "income").
		Select("COALESCE(SUM(amount), 0)").Scan(&income)

	config.DB.Model(&models.Record{}).
		Where("type = ?", "expense").
		Select("COALESCE(SUM(amount), 0)").Scan(&expense)

	// category-wise totals
	type CategoryTotal struct {
		Category string
		Total    float64
	}

	var categoryTotals []CategoryTotal

	config.DB.Model(&models.Record{}).
		Select("category, SUM(amount) as total").
		Group("category").
		Scan(&categoryTotals)

	// recent activity (last 5 records)
	var recent []models.Record
	config.DB.Order("date desc").Limit(5).Find(&recent)

	// monthly trends
	type MonthlyTrend struct {
		Month  string
		Total  float64
	}

	var monthly []MonthlyTrend

	config.DB.Model(&models.Record{}).
		Select("strftime('%Y-%m', date) as month, SUM(amount) as total").
		Group("month").
		Order("month asc").
		Scan(&monthly)

	return map[string]interface{}{
		"total_income":     income,
		"total_expense":    expense,
		"net_balance":      income - expense,
		"category_totals":  categoryTotals,
		"recent_activity":  recent,
		"monthly_trends":   monthly,
	}, nil
}