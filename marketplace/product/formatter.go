package product

import uuid "github.com/satori/go.uuid"

type ProductFormatter struct {
	ID     uuid.UUID
	Name   string
	Detail string
	Price  int
	Stock  int
}

func FormatProduct(product Products) ProductFormatter {
	productFormatter := ProductFormatter{}
	productFormatter.ID = product.ID
	productFormatter.Name = product.Name
	productFormatter.Detail = product.Detail
	productFormatter.Price = product.Price
	productFormatter.Stock = product.Stock

	return productFormatter
}
