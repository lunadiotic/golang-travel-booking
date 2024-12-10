// internal/usecase/user_usecase_test.go
package usecase_test

import (
	"errors"
	"testing"

	"github.com/lunadiotic/golang-travel-booking/internal/domain/entity"
	"github.com/lunadiotic/golang-travel-booking/internal/domain/repository/mocks"
	"github.com/lunadiotic/golang-travel-booking/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestUserUseCase_Register(t *testing.T) {
	// Shared test data
	const (
		validEmail    = "test@example.com"
		existingEmail = "existing@example.com"
		validPassword = "password123"
		fullName      = "Test User"
	)

	// Helper function untuk membuat user baru
	newTestUser := func(email, password, fullName string) *entity.User {
		return &entity.User{
			Email:    email,
			Password: password,
			FullName: fullName,
		}
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		user := newTestUser(validEmail, validPassword, fullName)

		mockRepo.On("FindByEmail", user.Email).Return(nil, nil)
		mockRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(nil)

		err := useCase.Register(user)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Email Already Exists", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		user := newTestUser(existingEmail, validPassword, fullName)
		existingUser := &entity.User{ID: "existing-id"}

		mockRepo.On("FindByEmail", user.Email).Return(existingUser, nil)

		err := useCase.Register(user)
		assert.Equal(t, usecase.ErrEmailAlreadyExists, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty Email", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		user := newTestUser("", validPassword, fullName)

		err := useCase.Register(user)
		assert.Equal(t, usecase.ErrInvalidInput, err)
	})

	t.Run("Empty Password", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		user := newTestUser(validEmail, "", fullName)

		err := useCase.Register(user)
		assert.Equal(t, usecase.ErrInvalidInput, err)
	})

	t.Run("Empty Full Name", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		user := newTestUser(validEmail, validPassword, "")

		err := useCase.Register(user)
		assert.Error(t, err)
		assert.Equal(t, usecase.ErrInvalidInput, err)
		// Tidak perlu AssertExpectations karena repository tidak dipanggil
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		user := newTestUser(validEmail, validPassword, fullName)
		dbError := errors.New("database error")

		mockRepo.On("FindByEmail", user.Email).Return(nil, nil)
		mockRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(dbError)

		err := useCase.Register(user)
		assert.Error(t, err)
		assert.Equal(t, dbError, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserUseCase_Login(t *testing.T) {
	// Shared test data
	const (
		validPassword   = "password123"
		invalidPassword = "wrongpassword"
		testEmail       = "test@example.com"
	)

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(validPassword), bcrypt.DefaultCost)
		storedUser := &entity.User{
			ID:       "user-123",
			Email:    testEmail,
			Password: string(hashedPassword),
		}

		mockRepo.On("FindByEmail", storedUser.Email).Return(storedUser, nil)

		token, user, err := useCase.Login(storedUser.Email, validPassword)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, storedUser.Email, user.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Password", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(validPassword), bcrypt.DefaultCost)
		storedUser := &entity.User{
			ID:       "user-123",
			Email:    testEmail,
			Password: string(hashedPassword),
		}

		mockRepo.On("FindByEmail", storedUser.Email).Return(storedUser, nil)

		token, user, err := useCase.Login(storedUser.Email, invalidPassword)
		assert.Error(t, err)
		assert.Equal(t, usecase.ErrInvalidCredentials, err)
		assert.Empty(t, token)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		mockRepo.On("FindByEmail", "nonexistent@example.com").Return(nil, nil)

		token, user, err := useCase.Login("nonexistent@example.com", validPassword)
		assert.Error(t, err)
		assert.Equal(t, usecase.ErrInvalidCredentials, err)
		assert.Empty(t, token)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty Email", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		token, user, err := useCase.Login("", validPassword)
		assert.Error(t, err)
		assert.Equal(t, usecase.ErrInvalidInput, err)
		assert.Empty(t, token)
		assert.Nil(t, user)
	})

	t.Run("Empty Password", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		token, user, err := useCase.Login(testEmail, "")
		assert.Error(t, err)
		assert.Equal(t, usecase.ErrInvalidInput, err)
		assert.Empty(t, token)
		assert.Nil(t, user)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		mockRepo.On("FindByEmail", testEmail).Return(nil, errors.New("database error"))

		token, user, err := useCase.Login(testEmail, validPassword)
		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Token Generation Error", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(validPassword), bcrypt.DefaultCost)
		storedUser := &entity.User{
			Email:    testEmail,
			Password: string(hashedPassword),
			// ID sengaja dikosongkan untuk memicu error token generation
		}

		mockRepo.On("FindByEmail", storedUser.Email).Return(storedUser, nil)

		token, user, err := useCase.Login(storedUser.Email, validPassword)
		assert.Error(t, err)
		assert.Equal(t, usecase.ErrTokenGeneration, err)
		assert.Empty(t, token)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})
}
