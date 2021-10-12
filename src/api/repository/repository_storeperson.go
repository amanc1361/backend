package repository

import "back-account/src/api/models"

type StorePerson interface {
	Save(models.StorePerson) (models.StorePerson, error)
	FindAll(companyid int, search string) ([]models.StorePerson, error)
	FindById(uint32) (models.StorePerson, error)
	Update(models.StorePerson) (int64, error)
	Delete(int32) (int64, error)
}
