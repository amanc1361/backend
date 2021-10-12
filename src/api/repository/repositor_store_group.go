package repository

import "back-account/src/api/models"

type StoreGroupRepository interface {
	Save(models.StoreGroup) (models.StoreGroup, error)
	FindAll(companyid int, search string) ([]models.StoreGroup, error)
	FindById(uint32) (models.StoreGroup, error)
	Update(models.StoreGroup) (int64, error)
	Delete(int32) (int64, error)
}
