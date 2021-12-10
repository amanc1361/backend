package repository

import (
	"back-account/src/api/models"
)

type DocumentMergeRepository interface {
	Save(models.DocumentMerge) (models.DocumentMerge, error)
	FindAll(companyid int, yearid int) ([]models.DocumentMerge, error)
	GetLastDocumetMergeCode(companyid int, yearid int) (int, error)
	GetLedgerDocument(docledger models.DocumentMerge) ([]models.LedgerDocument, error)
	Delete(int) (int64, error)
}
