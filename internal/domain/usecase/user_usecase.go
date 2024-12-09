package usecase

import "github.com/lunadiotic/golang-travel-booking/internal/domain/entity"

type UserUseCase interface {
	Register(user *entity.User) error
	Login(email, password string) (string, error) // returns JWT token
	GetProfile(id string) (*entity.User, error)
	UpdateProfile(user *entity.User) error
}
