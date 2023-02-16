package reposerviceUser

import (
	"errors"
	modelUser "project/model/user"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegistrasiUser(input modelUser.RegisterUserInput) (modelUser.User, error)
	Login(input modelUser.LoginInput) (modelUser.User, error)
	IsAvailableEmail(input modelUser.CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (modelUser.User, error)
	GetUserByID(ID int) (modelUser.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegistrasiUser(input modelUser.RegisterUserInput) (modelUser.User, error) {
	user := modelUser.User{}

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

func (s *service) Login(input modelUser.LoginInput) (modelUser.User, error) {
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

func (s *service) IsAvailableEmail(input modelUser.CheckEmailInput) (bool, error) {
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

func (s *service) SaveAvatar(ID int, fileLocation string) (modelUser.User, error) {
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

func (s *service) GetUserByID(ID int) (modelUser.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}
