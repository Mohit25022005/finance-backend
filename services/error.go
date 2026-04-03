package services

import "errors"

// Sentinel errors for consistent error handling across services.
// Use errors.Is() to compare these in controllers.
var (
	ErrUserNotFound   = errors.New("user not found")
	ErrRecordNotFound = errors.New("record not found")

	ErrInvalidRole   = errors.New("invalid role")
	ErrInvalidType   = errors.New("invalid record type")
	ErrInvalidAmount = errors.New("amount must be greater than zero")

	ErrMissingFields  = errors.New("required fields are missing")
	ErrDuplicateEmail = errors.New("email already exists")

	ErrUnauthorized = errors.New("unauthorized")
)