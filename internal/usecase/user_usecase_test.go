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
}

func TestUserUseCase_GetProfile(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		expectedUser := &entity.User{
			ID:       "user-123",
			Email:    "test@example.com",
			FullName: "Test User",
		}

		mockRepo.On("FindByID", expectedUser.ID).Return(expectedUser, nil)

		user, err := useCase.GetProfile(expectedUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty ID", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		user, err := useCase.GetProfile("")
		assert.Error(t, err)
		assert.Equal(t, usecase.ErrInvalidInput, err)
		assert.Nil(t, user)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		mockRepo.On("FindByID", "non-existent").Return(nil, nil)

		user, err := useCase.GetProfile("non-existent")
		assert.Error(t, err)
		assert.Equal(t, usecase.ErrUserNotFound, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		dbError := errors.New("database error")
		mockRepo.On("FindByID", "user-123").Return(nil, dbError)

		user, err := useCase.GetProfile("user-123")
		assert.Error(t, err)
		assert.Equal(t, dbError, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserUseCase_UpdateProfile(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		existingUser := &entity.User{
			ID:       "user-123",
			Email:    "test@example.com",
			FullName: "Old Name",
			Phone:    "123456",
		}

		updatedUser := &entity.User{
			ID:       "user-123",
			FullName: "New Name",
			Phone:    "654321",
		}

		mockRepo.On("FindByID", existingUser.ID).Return(existingUser, nil)
		mockRepo.On("Update", mock.AnythingOfType("*entity.User")).Return(nil)

		err := useCase.UpdateProfile(updatedUser)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Nil User", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		err := useCase.UpdateProfile(nil)
		assert.Error(t, err)
		assert.Equal(t, usecase.ErrInvalidInput, err)
	})

	t.Run("Empty ID", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		user := &entity.User{
			FullName: "Test User",
			Phone:    "123456",
		}

		err := useCase.UpdateProfile(user)
		assert.Error(t, err)
		assert.Equal(t, usecase.ErrInvalidInput, err)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		user := &entity.User{
			ID:       "non-existent",
			FullName: "Test User",
			Phone:    "123456",
		}

		mockRepo.On("FindByID", "non-existent").Return(nil, nil)

		err := useCase.UpdateProfile(user)
		assert.Error(t, err)
		assert.Equal(t, usecase.ErrUserNotFound, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error on FindByID", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		user := &entity.User{
			ID:       "user-123",
			FullName: "Test User",
			Phone:    "123456",
		}

		dbError := errors.New("database error")
		mockRepo.On("FindByID", user.ID).Return(nil, dbError)

		err := useCase.UpdateProfile(user)
		assert.Error(t, err)
		assert.Equal(t, dbError, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error on Update", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		useCase := usecase.NewUserUseCase(mockRepo)

		existingUser := &entity.User{
			ID:       "user-123",
			Email:    "test@example.com",
			FullName: "Old Name",
			Phone:    "123456",
		}

		updatedUser := &entity.User{
			ID:       "user-123",
			FullName: "New Name",
			Phone:    "654321",
		}

		dbError := errors.New("database error")
		mockRepo.On("FindByID", existingUser.ID).Return(existingUser, nil)
		mockRepo.On("Update", mock.AnythingOfType("*entity.User")).Return(dbError)

		err := useCase.UpdateProfile(updatedUser)
		assert.Error(t, err)
		assert.Equal(t, dbError, err)
		mockRepo.AssertExpectations(t)
	})
}
