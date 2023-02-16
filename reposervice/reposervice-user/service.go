package reposerviceuser

import (
	"errors"
	"project/model/user"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegistrasiUser(input user.RegisterUserInput) (user.User, error)
	Login(input user.LoginInput) (user.User, error)
	IsAvailableEmail(input user.CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (user.User, error)
	GetUserByID(ID int) (user.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegistrasiUser(input user.RegisterUserInput) (user.User, error) {
	user := user.User{}

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

func (s *service) Login(input user.LoginInput) (user.User, error) {
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

func (s *service) IsAvailableEmail(input user.CheckEmailInput) (bool, error) {
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

func (s *service) SaveAvatar(ID int, fileLocation string) (user.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updateUser, err := s.repository.Update(user)
	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}

func (s *service) GetUserByID(ID int) (user.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}
