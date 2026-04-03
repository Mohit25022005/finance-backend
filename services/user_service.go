package services

import (
	"fmt"

	"finance-backend/config"
	"finance-backend/models"
	"finance-backend/repository"
)

// UserService defines the business logic contract for users.
type UserService interface {
	Create(req models.CreateUserRequest) (*models.User, error)
	GetAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	Update(id uint, req models.UpdateUserRequest) (*models.User, error)
	Delete(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

// NewUserService creates a UserService with the given repository.
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(req models.CreateUserRequest) (*models.User, error) {
	if !req.Role.IsValid() {
		return nil, ErrInvalidRole
	}

	// Check for duplicate email
	existing, err := s.repo.GetByEmail(req.Email)
	if err != nil && !config.IsNotFound(err) {
		return nil, fmt.Errorf("check existing email: %w", err)
	}
	if existing != nil {
		return nil, ErrDuplicateEmail
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
		IsActive: true,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}

func (s *userService) GetAll() ([]models.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("get all users: %w", err)
	}
	return users, nil
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		if config.IsNotFound(err) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("get user by id: %w", err)
	}
	return user, nil
}

func (s *userService) Update(id uint, req models.UpdateUserRequest) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		if config.IsNotFound(err) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("fetch user for update: %w", err)
	}

	// Apply only provided fields
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Role != nil {
		if !req.Role.IsValid() {
			return nil, ErrInvalidRole
		}
		user.Role = *req.Role
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	if err := s.repo.Update(user); err != nil {
		if config.IsNotFound(err) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("update user: %w", err)
	}

	return user, nil
}

func (s *userService) Delete(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		if config.IsNotFound(err) {
			return ErrUserNotFound
		}
		return fmt.Errorf("fetch user for delete: %w", err)
	}

	if err := s.repo.Delete(id); err != nil {
		if config.IsNotFound(err) {
			return ErrUserNotFound
		}
		return fmt.Errorf("delete user: %w", err)
	}

	return nil
}