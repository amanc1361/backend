package repository

import "back-account/src/api/models"

type StroeObjectRepository interface {
	Save(models.StoreObject) (models.StoreObject, error)
	FindAll() ([]models.StoreObject, error)
	GetList(companyid int, storeid int) ([]models.PublicStoreObject, error)
	FindById(uint32) (models.StoreObject, error)
	Update(models.StoreObject) (int64, error)
	Delete(int32) (int64, error)
}
