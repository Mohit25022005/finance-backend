package services

import (
	
	"time"

	"finance-backend/models"
	"gorm.io/gorm"
)

type DashboardService interface {
	GetDashboardData(userID uint) (*DashboardData, error)
}

type DashboardData struct {
	TotalIncome    float64                 `json:"total_income"`
	TotalExpense   float64                 `json:"total_expense"`
	NetBalance     float64                 `json:"net_balance"`
	CategoryTotals []CategoryTotal         `json:"category_totals"`
	RecentActivity []models.RecordResponse `json:"recent_activity"`
	MonthlyTrends  []MonthlyTrend          `json:"monthly_trends"`
}

type CategoryTotal struct {
	Category string  `json:"category"`
	Total    float64 `json:"total"`
}

type MonthlyTrend struct {
	Month string  `json:"month"`
	Type  string  `json:"type"`
	Total float64 `json:"total"`
}

type dashboardService struct {
	db *gorm.DB
}

func NewDashboardService(db *gorm.DB) DashboardService {
	return &dashboardService{db: db}
}

func (s *dashboardService) GetDashboardData(userID uint) (*DashboardData, error) {
	income, err := s.sumByType(userID, models.RecordTypeIncome)
	if err != nil {
		return nil, err
	}

	expense, err := s.sumByType(userID, models.RecordTypeExpense)
	if err != nil {
		return nil, err
	}

	categoryTotals, err := s.getCategoryTotals(userID)
	if err != nil {
		return nil, err
	}

	recent, err := s.getRecentActivity(userID)
	if err != nil {
		return nil, err
	}

	monthly, err := s.getMonthlyTrends(userID)
	if err != nil {
		return nil, err
	}

	return &DashboardData{
		TotalIncome:    income,
		TotalExpense:   expense,
		NetBalance:     income - expense,
		CategoryTotals: categoryTotals,
		RecentActivity: recent,
		MonthlyTrends:  monthly,
	}, nil
}

func (s *dashboardService) sumByType(userID uint, t models.RecordType) (float64, error) {
	var total float64
	err := s.db.Model(&models.Record{}).
		Where("user_id = ? AND type = ?", userID, t).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&total).Error
	return total, err
}

func (s *dashboardService) getCategoryTotals(userID uint) ([]CategoryTotal, error) {
	var totals []CategoryTotal
	err := s.db.Model(&models.Record{}).
		Where("user_id = ?", userID).
		Select("category, SUM(amount) as total").
		Group("category").
		Order("total DESC").
		Scan(&totals).Error
	return totals, err
}

func (s *dashboardService) getRecentActivity(userID uint) ([]models.RecordResponse, error) {
	var records []models.Record
	err := s.db.Where("user_id = ?", userID).
		Order("date DESC").
		Limit(5).
		Find(&records).Error

	if err != nil {
		return nil, err
	}

	res := make([]models.RecordResponse, len(records))
	for i, r := range records {
		res[i] = r.ToResponse()
	}
	return res, nil
}

func (s *dashboardService) getMonthlyTrends(userID uint) ([]MonthlyTrend, error) {
	var trends []MonthlyTrend
	start := time.Now().AddDate(0, -12, 0)

	err := s.db.Model(&models.Record{}).
		Where("user_id = ? AND date >= ?", userID, start).
		Select("strftime('%Y-%m', date) as month, type, SUM(amount) as total").
		Group("month, type").
		Order("month ASC").
		Scan(&trends).Error

	return trends, err
}