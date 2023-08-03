package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
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
