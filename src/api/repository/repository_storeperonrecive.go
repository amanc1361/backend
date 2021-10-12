package repository

import "back-account/src/api/models"

type StorePersonRecive interface {
	Save(models.StorePersonRecive) (models.StorePersonRecive, error)
	FindAll(companyid int, search string) ([]models.StorePersonRecive, error)
	FindById(uint32) (models.StorePersonRecive, error)
	Update(models.StorePersonRecive) (int64, error)
	Delete(int32) (int64, error)
}
