package repository

import "back-account/src/api/models"

type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindById(uint32) (models.User, error)
	Update(int32, models.User) (int64, error)
	Delete(int32) (int64, error)
}
