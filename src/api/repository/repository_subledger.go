package repository

import (
	"back-account/src/api/models"
	"back-account/src/api/modelsout"
)

type SubLedgerRepository interface {
	Save(models.SubLedger) (modelsout.SubledgerOut, error)
	FindAll(companyid int) ([]modelsout.SubledgerOut, error)
	FindById(uint32) (modelsout.SubledgerOut, error)
	FindByCode(code uint, companyid int) (modelsout.SubledgerOut, error)
	Update(models.SubLedger) (int64, error)
	Delete(int32) (int64, error)
	GetLastCode(code uint, companyid int) (uint, error)
	GetByLedgerID(id uint) ([]modelsout.SubledgerOut, error)
}
