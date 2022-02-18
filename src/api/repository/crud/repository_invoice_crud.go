package crud

import (
	"back-account/src/api/models"

	"gorm.io/gorm"
)

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvocieRepository(db *gorm.DB) *invoiceRepository {
	return &invoiceRepository{db}
}

func (r *invoiceRepository) Save(invoice models.Invocie) (models.Invocie,error) {
//todo 
return models.Invocie{},nil

}