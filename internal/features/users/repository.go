package users

import "sprig/internal/model"

type UserRepository interface {
	AddUser(user *model.User) error
	FindByID(id uint64) (*model.User, error)
	Save(user *model.User) error
	DeleteByID(id uint64) error
}
