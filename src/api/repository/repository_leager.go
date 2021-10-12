package repository

import "back-account/src/api/models"

type LedgerRepository interface {
	Save(models.Ledger) (models.Custledger, error)
	FindAll(companyid int, search string) ([]models.Custledger, error)
	FindById(uint32) (models.Custledger, error)
	FindByCode(code uint, companyid int) (models.Custledger, error)
	GetByGroupID(id uint) ([]models.Custledger, error)
	Update(models.Ledger) (int64, error)
	Delete(int32) (int64, error)
	GetLastCode(code uint, companyid int) (uint, error)
}
