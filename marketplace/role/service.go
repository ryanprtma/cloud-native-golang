package role

import (
	"errors"
	"marketplace/user"

	uuid "github.com/satori/go.uuid"
)

type Service interface {
	GetMerchantByUserID(ID string) (Merchant, error)
	SaveRole(user user.User, input user.RegisterUserInput)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetMerchantByUserID(ID string) (Merchant, error) {
	merchant, err := s.repository.FindByID(ID)
	if err != nil {
		return merchant, err
	}

	if merchant.ID == uuid.Nil {
		return merchant, errors.New("user not found")
	}

	return merchant, nil
}

func (s *service) SaveRole(user user.User, input user.RegisterUserInput) {
	uuid := uuid.NewV4()
	roleId := input.RoleId
	if roleId == 0 {
		customer := Customer{}
		customer.ID = uuid
		customer.Address = "Known Here"
		customer.User = user
		customer.RoleId = input.RoleId
		s.repository.SaveCustomer(customer)
	} else if roleId == 1 {
		merchant := Merchant{}
		merchant.ID = uuid
		merchant.MerchantName = "test1"
		merchant.Address = "Know Here"
		merchant.User = user
		merchant.RoleId = input.RoleId
		s.repository.SaveMerchant(merchant)
	}
}
