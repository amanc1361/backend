package repository

import "back-account/src/api/models"

type CompanyTypesRepository interface {
	FindAll() ([]models.CompanyType, error)
	FindById(uint32) (models.CompanyType, error)
}
