// internal/domain/repository/user_repository.go
package repository

import (
	"github.com/lunadiotic/golang-travel-booking/internal/domain/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByID(id string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id string) error
}
