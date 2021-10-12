package repository

import "back-account/src/api/models"

type YearRepository interface {
	Save(models.Year) (models.Year, error)
	FindAll(companyid int) ([]models.Year, error)
	FindById(uint32) (models.Year, error)
	Update(models.Year) (int64, error)
	Delete(uint32) (int64, error)
}
