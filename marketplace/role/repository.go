package role

import (
	"gorm.io/gorm"
)

type Repository interface {
	SaveMerchant(merchant Merchant) (Merchant, error)
	SaveCustomer(customer Customer) (Customer, error)
	FindByID(ID string) (Merchant, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveMerchant(merchant Merchant) (Merchant, error) {
	err := r.db.Create(&merchant).Error
	if err != nil {
		return merchant, err
	}

	return merchant, nil
}

func (r *repository) SaveCustomer(customer Customer) (Customer, error) {
	err := r.db.Create(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *repository) FindByID(ID string) (Merchant, error) {
	var merchant Merchant
	err := r.db.Where("user_id = ?", ID).Find(&merchant).Error
	if err != nil {
		return merchant, err
	}
	return merchant, nil
}
