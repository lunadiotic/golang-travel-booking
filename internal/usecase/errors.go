package usecase

import "errors"

var (
	ErrEmailAlreadyExists  = errors.New("email already exists")
	ErrInvalidCredentials  = errors.New("invalid email or password")
	ErrInvalidInput        = errors.New("invalid input")
	ErrTokenGeneration     = errors.New("failed to generate token")
	ErrUserNotFound        = errors.New("user not found")
	ErrDatabaseError       = errors.New("database error occurred")
	ErrDestinationNotFound = errors.New("destination not found")
)
