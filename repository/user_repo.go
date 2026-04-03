package repository

import (
	"fmt"

	"finance-backend/models"

	"gorm.io/gorm"
)

// UserRepository defines the contract for user data access.
// Using an interface lets you swap implementations (e.g. mock in tests).
type UserRepository interface {
	Create(user *models.User) error
	GetAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository with the given DB instance.
// Accepts *gorm.DB instead of using the global config.DB directly.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("get all users: %w", err)
	}
	return users, nil
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, fmt.Errorf("get user by id %d: %w", id, err)
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("get user by email %s: %w", email, err)
	}
	return &user, nil
}

func (r *userRepository) Update(user *models.User) error {
	// Save() updates all fields including zero values — use Updates() with
	// a map/struct to only update fields that were actually provided.
	result := r.db.Model(user).Updates(user)
	if result.Error != nil {
		return fmt.Errorf("update user %d: %w", user.ID, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("update user %d: %w", user.ID, gorm.ErrRecordNotFound)
	}
	return nil
}

func (r *userRepository) Delete(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return fmt.Errorf("delete user %d: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("delete user %d: %w", id, gorm.ErrRecordNotFound)
	}
	return nil
}