package usecase

import (
	"github.com/lunadiotic/golang-travel-booking/internal/domain/entity"
	"github.com/lunadiotic/golang-travel-booking/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *userUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) Register(user *entity.User) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Check if email already exists
	existingUser, err := u.userRepo.FindByEmail(user.Email)
	if err == nil && existingUser != nil {
		return ErrEmailAlreadyExists
	}

	return u.userRepo.Create(user)
}

func (u *userUseCase) Login(email, password string) (string, error) {
	// Implementasi login
	return "", nil
}

func (u *userUseCase) GetProfile(id string) (*entity.User, error) {
	// Implementasi get profile
	return u.userRepo.FindByID(id)
}

func (u *userUseCase) UpdateProfile(user *entity.User) error {
	// Implementasi update profile
	return u.userRepo.Update(user)
}
