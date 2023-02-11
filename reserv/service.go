package reserv

import (
	"errors"
	"project/model"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegistrasiUser(input model.RegisterUserInput) (model.User, error)
	Login(input model.LoginInput) (model.User, error)
	IsAvailableEmail(input model.CheckEmailInput) (bool, error)
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

func (s *service) Login(input model.LoginInput) (model.User, error) {
	email := input.Email
	password := input.Password

	newUser, err := s.repository.FindByEmail(email)
	if err != nil {
		return newUser, err
	}

	if newUser.ID == 0 {
		return newUser, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(password))
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) IsAvailableEmail(input model.CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}
