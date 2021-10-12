package repository

import (
	"back-account/src/api/models"
	"back-account/src/api/modelsout"
)

type DetailedRepository interface {
	Save(models.Detailed) (models.Detailed, error)
	FindAll(search string, companyid int) ([]models.Detailed, uint, error)
	FindById(uint32) (models.Detailed, error)
	FindByCode(uint, int) (models.Detailed, error)
	Update(int32, models.Detailed) (int64, error)
	Delete(int32) (int64, error)
	GetLastCode(int) (uint, error)
	GetFlow(yearid int, detaildid int) ([]modelsout.DetailedFlow, error)
}
