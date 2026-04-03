package models

import (
	"time"

	"gorm.io/gorm"
)

// Role is a typed string to prevent arbitrary role values being assigned.
type Role string

const (
	RoleAdmin   Role = "admin"
	RoleAnalyst Role = "analyst"
	RoleViewer  Role = "viewer"
)

// IsValid checks if a role value is one of the allowed roles.
func (r Role) IsValid() bool {
	switch r {
	case RoleAdmin, RoleAnalyst, RoleViewer:
		return true
	}
	return false
}

// String implements the Stringer interface.
func (r Role) String() string {
	return string(r)
}

// User represents a system user with role-based access.
type User struct {
	ID        uint           `json:"id"         gorm:"primaryKey"`
	Name      string         `json:"name"       gorm:"not null"`
	Email     string         `json:"email"      gorm:"uniqueIndex;not null"`
	Role      Role           `json:"role"       gorm:"not null;default:viewer"`
	IsActive  bool           `json:"is_active"  gorm:"not null;default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index" swaggertype:"string"`
}

// UserResponse is what gets sent over the API — never expose the full model directly.
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      Role      `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateUserRequest is the validated input shape for creating a user.
type CreateUserRequest struct {
	Name  string `json:"name"  binding:"required,min=2,max=100"`
	Email string `json:"email" binding:"required,email"`
	Role  Role   `json:"role"  binding:"required"`
}

// UpdateUserRequest is the validated input shape for updating a user.
type UpdateUserRequest struct {
	Name     *string `json:"name"      binding:"omitempty,min=2,max=100"`
	Role     *Role   `json:"role"      binding:"omitempty"`
	IsActive *bool   `json:"is_active" binding:"omitempty"`
}

// ToResponse converts a User model to a safe API response.
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
	}
}

// TableName explicitly sets the table name to avoid GORM pluralization surprises.
func (User) TableName() string {
	return "users"
}