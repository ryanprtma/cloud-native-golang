package product

import (
	uuid "github.com/satori/go.uuid"
)

type Service interface {
	CreateProduct(input CreateProductInput) (Products, error)
	UpdateProductTransaction(input CreateProductInput, id string) (Products, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateProduct(input CreateProductInput) (Products, error) {

	uuid := uuid.NewV4()
	product := Products{}
	product.ID = uuid
	product.Name = input.Name
	product.Detail = input.Detail
	product.Price = input.Price
	product.Stock = input.Stock
	product.Merchant = input.Merchant

	newProduct, err := s.repository.Save(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) UpdateProductTransaction(input CreateProductInput, id string) (Products, error) {

	product, err := s.repository.FindByID(id)
	if err != nil {
		return product, err
	}

	qty := product.Stock - input.Stock
	updateProducts, err := s.repository.Update(id, qty)
	if err != nil {
		return updateProducts, err
	}

	return updateProducts, nil
}
