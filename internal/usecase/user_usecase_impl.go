package usecase

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/lunadiotic/golang-travel-booking/internal/domain/entity"
	"github.com/lunadiotic/golang-travel-booking/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo  repository.UserRepository
	jwtSecret string // tambahkan ini
}

func NewUserUseCase(userRepo repository.UserRepository, jwtSecret string) *userUseCase {
	return &userUseCase{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (u *userUseCase) Register(user *entity.User) error {
	// Validasi input
	if user.Email == "" || user.Password == "" || user.FullName == "" {
		return ErrInvalidInput
	}

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

func (u *userUseCase) generateToken(user *entity.User) (string, error) {
	if user == nil {
		return "", ErrTokenGeneration
	}

	// Validasi ID user
	if user.ID == "" {
		return "", ErrTokenGeneration
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return "", ErrTokenGeneration
	}

	return tokenString, nil
}

func (u *userUseCase) Login(email, password string) (string, *entity.User, error) {
	if email == "" || password == "" {
		return "", nil, ErrInvalidInput
	}

	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", nil, err
	}

	if user == nil {
		return "", nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, ErrInvalidCredentials
	}

	token, err := u.generateToken(user)
	if err != nil {
		return "", nil, err // Gunakan error dari generateToken
	}

	return token, user, nil
}

func (u *userUseCase) GetProfile(id string) (*entity.User, error) {
	if id == "" {
		return nil, ErrInvalidInput
	}

	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (u *userUseCase) UpdateProfile(user *entity.User) error {
	if user == nil || user.ID == "" {
		return ErrInvalidInput
	}

	existingUser, err := u.userRepo.FindByID(user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return ErrUserNotFound
	}

	// Update only allowed fields
	existingUser.FullName = user.FullName
	existingUser.Phone = user.Phone

	return u.userRepo.Update(existingUser)
}
