package repository

import "back-account/src/api/models"

type CompanyRepository interface {
	Save(models.Company) (models.Company, error)
	FindAll(userid int) ([]models.Company, error)
	FindById(uint32) (models.Company, error)
	Update(int32, models.Company) (int64, error)
	Delete(int32) (int64, error)
}
