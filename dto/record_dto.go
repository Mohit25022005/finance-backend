package dto

type CreateRecordDTO struct {
	Amount   float64 `json:"amount" binding:"required"`
	Type     string  `json:"type" binding:"required"`
	Category string  `json:"category"`
	Notes    string  `json:"notes"`
	UserID   uint    `json:"user_id"`
}