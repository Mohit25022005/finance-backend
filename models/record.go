package models

import (
	"time"

	"gorm.io/gorm"
)

// RecordType is a typed string to prevent arbitrary transaction types.
type RecordType string

const (
	RecordTypeIncome  RecordType = "income"
	RecordTypeExpense RecordType = "expense"
)

// IsValid checks if a record type is one of the allowed values.
func (rt RecordType) IsValid() bool {
	switch rt {
	case RecordTypeIncome, RecordTypeExpense:
		return true
	}
	return false
}

// String implements the Stringer interface.
func (rt RecordType) String() string {
	return string(rt)
}

// Record represents a single financial transaction.
type Record struct {
	ID        uint           `json:"id"         gorm:"primaryKey"`
	Amount    float64        `json:"amount"     gorm:"not null"`
	Type      RecordType     `json:"type"       gorm:"not null;index"`
	Category  string         `json:"category"   gorm:"not null;index"`
	Date      time.Time      `json:"date"       gorm:"not null;index"`
	Notes     string         `json:"notes"`
	UserID    uint           `json:"user_id"    gorm:"not null;index"`
	User      User           `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index" swaggertype:"string"`
}

// RecordResponse is what gets sent over the API.
type RecordResponse struct {
	ID        uint       `json:"id"`
	Amount    float64    `json:"amount"`
	Type      RecordType `json:"type"`
	Category  string     `json:"category"`
	Date      time.Time  `json:"date"`
	Notes     string     `json:"notes"`
	UserID    uint       `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// CreateRecordRequest is the validated input shape for creating a record.
type CreateRecordRequest struct {
	Amount   float64    `json:"amount"   binding:"required,gt=0"`
	Type     RecordType `json:"type"     binding:"required"`
	Category string     `json:"category" binding:"required,min=2,max=100"`
	Date     time.Time  `json:"date"     binding:"required"`
	Notes    string     `json:"notes"    binding:"omitempty,max=500"`
}

// UpdateRecordRequest is the validated input shape for updating a record.
// All fields are pointers so we can distinguish "not provided" from "zero value".
type UpdateRecordRequest struct {
	Amount   *float64    `json:"amount"   binding:"omitempty,gt=0"`
	Type     *RecordType `json:"type"     binding:"omitempty"`
	Category *string     `json:"category" binding:"omitempty,min=2,max=100"`
	Date     *time.Time  `json:"date"     binding:"omitempty"`
	Notes    *string     `json:"notes"    binding:"omitempty,max=500"`
}

// RecordFilterParams holds validated query parameters for listing records.
type RecordFilterParams struct {
	Type     RecordType `form:"type"     binding:"omitempty"`
	Category string     `form:"category" binding:"omitempty"`
	Search   string     `form:"search"   binding:"omitempty,max=100"`
	Page     int        `form:"page"     binding:"omitempty,min=1"`
	Limit    int        `form:"limit"    binding:"omitempty,min=1,max=100"`
}

// Normalize sets default pagination values if not provided.
func (f *RecordFilterParams) Normalize() {
	if f.Page == 0 {
		f.Page = 1
	}
	if f.Limit == 0 {
		f.Limit = 10
	}
}

// Offset calculates the DB offset from page and limit.
func (f *RecordFilterParams) Offset() int {
	return (f.Page - 1) * f.Limit
}

// ToResponse converts a Record model to a safe API response.
func (r *Record) ToResponse() RecordResponse {
	return RecordResponse{
		ID:        r.ID,
		Amount:    r.Amount,
		Type:      r.Type,
		Category:  r.Category,
		Date:      r.Date,
		Notes:     r.Notes,
		UserID:    r.UserID,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

// TableName explicitly sets the table name.
func (Record) TableName() string {
	return "records"
}