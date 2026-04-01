package models

import "time"

type Record struct {
	ID       uint      `gorm:"primaryKey"`
	Amount   float64   `json:"amount"`
	Type     string    `json:"type"` // income/expense
	Category string    `json:"category"`
	Date     time.Time `json:"date"`
	Notes    string    `json:"notes"`
	UserID   uint      `json:"user_id"`
}