package reserv

import (
	"project/model"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegistrasiUser(input model.RegisterUserInput) (model.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegistrasiUser(input model.RegisterUserInput) (model.User, error) {
	user := model.User{}

	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email
	Password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(Password)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}

	return newUser, nil
}
