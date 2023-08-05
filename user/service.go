package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterUserInput) (User, error)
	Login(input LoginUserInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(input RegisterUserInput) (User, error) {
	user := User{}

	user.Name = input.Name
	user.Nim = input.Nim
	user.Email = input.Email
	user.Division = input.Division
	user.NoHP = input.NoHP
	user.AlasanDaftar = input.AlasanDaftar
	passworHash, err := bcrypt.GenerateFromPassword([]byte(input.Nim), bcrypt.DefaultCost)
	user.Password = string(passworHash)
	user.RoleID = false

	registeredUser, err := s.repository.Save(user)
	if err != nil {
		return registeredUser, err
	}

	return registeredUser, nil
}

func (s *service) Login(input LoginUserInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Email not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
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
