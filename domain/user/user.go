package user

import "github.com/BrandokVargas/api-back-dportinsight/model"

type UseCase interface {
	RegisterUser(*model.User) error
	GetAllUsers() (model.Users, error)
}

type Repository interface {
	RegisterUser(*model.User) error
	GetAllUsers() (model.Users, error)
}
