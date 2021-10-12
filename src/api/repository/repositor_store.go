package repository

import "back-account/src/api/models"

type StoreRepository interface {
	Save(models.Store) (models.Store, error)
	FindAll(companyid int, search string) ([]models.Store, error)
	FindById(uint32) (models.Store, error)
	Update(models.Store) (int64, error)
	Delete(int32) (int64, error)
}
