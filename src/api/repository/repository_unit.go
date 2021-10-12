package repository

import "back-account/src/api/models"

type UnitRepository interface {
	Save(models.Unit) (models.Unit, error)
	FindAll(companyid int, search string) ([]models.Unit, error)
	FindById(uint32) (models.Unit, error)
	Update(models.Unit) (int64, error)
	Delete(int32) (int64, error)
}
