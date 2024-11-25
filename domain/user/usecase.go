package user

import (
	"fmt"
	"time"

	"github.com/BrandokVargas/api-back-dportinsight/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repository Repository
}

func NewUser(repository Repository) *User {
	return &User{repository: repository}
}

func (u User) RegisterUser(m *model.User) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	m.ID = ID

	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("%s %w", "bcrypt.GenerateFromPassword()", err)
	}
	m.Password = string(password)
	m.CreatedAt = time.Now().Unix()

	err = u.repository.RegisterUser(m)

	if err != nil {
		return fmt.Errorf("%s %w", "repository.RegisterUser()", err)
	}

	m.Password = ""
	return nil
}

func (u User) GetAllUsers() (model.Users, error) {
	users, err := u.repository.GetAllUsers()

	if err != nil {
		return nil, fmt.Errorf("%s %w", "repository.GetAllUsers()", err)
	}
	return users, nil
}
