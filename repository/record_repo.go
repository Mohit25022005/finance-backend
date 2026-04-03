package repository

import (
	"fmt"

	"finance-backend/models"

	"gorm.io/gorm"
)

// RecordRepository defines the contract for record data access.
type RecordRepository interface {
	Create(record *models.Record) error
	GetByID(id uint) (*models.Record, error)
	Update(record *models.Record) error
	Delete(id uint) error
	GetFiltered(params models.RecordFilterParams) ([]models.Record, int64, error)
}

type recordRepository struct {
	db *gorm.DB
}

// NewRecordRepository creates a new RecordRepository with the given DB instance.
func NewRecordRepository(db *gorm.DB) RecordRepository {
	return &recordRepository{db: db}
}

func (r *recordRepository) Create(record *models.Record) error {
	if err := r.db.Create(record).Error; err != nil {
		return fmt.Errorf("create record: %w", err)
	}
	return nil
}

func (r *recordRepository) GetByID(id uint) (*models.Record, error) {
	var record models.Record
	err := r.db.First(&record, id).Error
	if err != nil {
		return nil, fmt.Errorf("get record by id %d: %w", id, err)
	}
	return &record, nil
}

func (r *recordRepository) Update(record *models.Record) error {
	result := r.db.Model(record).Updates(record)
	if result.Error != nil {
		return fmt.Errorf("update record %d: %w", record.ID, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("update record %d: %w", record.ID, gorm.ErrRecordNotFound)
	}
	return nil
}

func (r *recordRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Record{}, id)
	if result.Error != nil {
		return fmt.Errorf("delete record %d: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("delete record %d: %w", id, gorm.ErrRecordNotFound)
	}
	return nil
}

// GetFiltered returns paginated, filtered records and the total count.
// Returning total count lets the client build pagination controls.
func (r *recordRepository) GetFiltered(params models.RecordFilterParams) ([]models.Record, int64, error) {
	params.Normalize()

	query := r.buildFilterQuery(params)

	// Count total matching records BEFORE applying pagination.
	// This gives the client the total pages count.
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count filtered records: %w", err)
	}

	var records []models.Record
	err := query.
		Order("date DESC").
		Limit(params.Limit).
		Offset(params.Offset()).
		Find(&records).Error
	if err != nil {
		return nil, 0, fmt.Errorf("get filtered records: %w", err)
	}

	return records, total, nil
}

// buildFilterQuery constructs the WHERE clause from filter params.
// Extracted as a separate method to keep GetFiltered readable and allow
// reuse for the count query.
func (r *recordRepository) buildFilterQuery(params models.RecordFilterParams) *gorm.DB {
	query := r.db.Model(&models.Record{})

	if params.Type.IsValid() {
		query = query.Where("type = ?", params.Type)
	}
	if params.Category != "" {
		query = query.Where("category = ?", params.Category)
	}
	if params.Search != "" {
		search := "%" + params.Search + "%"
		query = query.Where("category LIKE ? OR notes LIKE ?", search, search)
	}

	return query
}

// GetDB exposes the underlying db — useful for dashboard/analytics queries
// that don't fit neatly into the standard CRUD interface.
func (r *recordRepository) GetDB() *gorm.DB {
	return r.db
}