package user

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error) //parameter interface mengassign fungsi yang diperlukan
	Login(input LoginInput) (User, error)
	CheckEmailAvailiblelity(input CheckEmailInput) (bool, error)
	SaveAvatar(ID string, fileLocation string) (User, error)
	GetUserByID(ID string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	uuid := uuid.NewV4()
	user.ID = uuid
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.RoleId = input.RoleId
	HashPassword, err := bcrypt.GenerateFromPassword([]byte(input.HashPassword), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.HashPassword = string(HashPassword)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == uuid.Nil {
		return user, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) CheckEmailAvailiblelity(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == uuid.Nil {
		return true, nil
	}

	return false, nil

}

func (s *service) SaveAvatar(ID string, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updadatesUser, err := s.repository.Update(user)

	if err != nil {
		return updadatesUser, err
	}

	return updadatesUser, nil

}

func (s *service) GetUserByID(ID string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == uuid.Nil {
		return user, errors.New("user not found")
	}

	return user, nil
}
