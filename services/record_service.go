package services

import (
	"fmt"

	"finance-backend/config"
	"finance-backend/models"
	"finance-backend/repository"
)

// RecordService defines the business logic contract for records.
type RecordService interface {
	Create(userID uint, req models.CreateRecordRequest) (*models.Record, error)
	GetFiltered(params models.RecordFilterParams) ([]models.Record, int64, error)
	GetByID(id uint) (*models.Record, error)
	Update(id uint, userID uint, req models.UpdateRecordRequest) (*models.Record, error)
	Delete(id uint, userID uint) error
}

type recordService struct {
	repo repository.RecordRepository
}

// NewRecordService creates a RecordService with the given repository.
func NewRecordService(repo repository.RecordRepository) RecordService {
	return &recordService{repo: repo}
}

func (s *recordService) Create(userID uint, req models.CreateRecordRequest) (*models.Record, error) {
	if !req.Type.IsValid() {
		return nil, ErrInvalidType
	}

	record := &models.Record{
		Amount:   req.Amount,
		Type:     req.Type,
		Category: req.Category,
		Date:     req.Date,
		Notes:    req.Notes,
		UserID:   userID,
	}

	if err := s.repo.Create(record); err != nil {
		return nil, fmt.Errorf("create record: %w", err)
	}

	return record, nil
}

func (s *recordService) GetFiltered(params models.RecordFilterParams) ([]models.Record, int64, error) {
	records, total, err := s.repo.GetFiltered(params)
	if err != nil {
		return nil, 0, fmt.Errorf("get filtered records: %w", err)
	}
	return records, total, nil
}

func (s *recordService) GetByID(id uint) (*models.Record, error) {
	record, err := s.repo.GetByID(id)
	if err != nil {
		if config.IsNotFound(err) {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("get record by id: %w", err)
	}
	return record, nil
}

func (s *recordService) Update(id uint, userID uint, req models.UpdateRecordRequest) (*models.Record, error) {
	record, err := s.repo.GetByID(id)
	if err != nil {
		if config.IsNotFound(err) {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("fetch record for update: %w", err)
	}

	//  ensure user owns the record
	if record.UserID != userID {
		return nil, fmt.Errorf("unauthorized")
	}

	// Apply only provided fields
	if req.Amount != nil {
		record.Amount = *req.Amount
	}
	if req.Type != nil {
		if !req.Type.IsValid() {
			return nil, ErrInvalidType
		}
		record.Type = *req.Type
	}
	if req.Category != nil {
		record.Category = *req.Category
	}
	if req.Date != nil {
		record.Date = *req.Date
	}
	if req.Notes != nil {
		record.Notes = *req.Notes
	}

	if err := s.repo.Update(record); err != nil {
		return nil, fmt.Errorf("update record: %w", err)
	}

	return record, nil
}

func (s *recordService) Delete(id uint, userID uint) error {
	record, err := s.repo.GetByID(id)
	if err != nil {
		if config.IsNotFound(err) {
			return ErrRecordNotFound
		}
		return fmt.Errorf("fetch record for delete: %w", err)
	}

	//  ensure user owns the record
	if record.UserID != userID {
		return fmt.Errorf("unauthorized")
	}

	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("delete record: %w", err)
	}

	return nil
}