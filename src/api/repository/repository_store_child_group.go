package repository

import "back-account/src/api/models"

type StoreChildSubGroupRepository interface {
	Save(models.StoreChildSubGroup) (models.StoreChildSubGroup, error)
	FindAll() ([]models.StoreChildSubGroup, error)
	FindById(uint32) (models.StoreChildSubGroup, error)
	Update(models.StoreChildSubGroup) (int64, error)
	Delete(int32) (int64, error)
}
