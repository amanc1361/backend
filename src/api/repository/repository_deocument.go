package repository

import (
	"back-account/src/api/models"
	"back-account/src/api/modelsout"
	"back-account/src/api/utils/pagination"
)

type DocumentRepository interface {
	Save(models.Document) (models.Document, error)
	FindAll(string, uint, int, int, string) (pagination.Pagination, error)
	FindByID(uint) (models.Document, error)
	FindByDocumentCode(uint, int, int) (models.Document, error)
	FindByDocumentDescription(string, int, int) ([]models.Document, error)
	FindRowsByDcoumentid(uint) ([]models.DocumentRows, error)
	Update(uint, int, models.Document) (uint, error)
	GetLastCode(uint, uint) (uint, error)
	Delete(uint) (int64, error)
	GetInfo(uint, uint) (modelsout.DocInfo, error)
	GetByType(companyid int, yearid int, doctype int) ([]models.Document, error)
	SortDoc(companyid int, yearid int) (uint, error)
	GetFirstDate(companyid int, yearid int) (string, error)
	GetLastDate(companyid int, yearid int) (string, error)
}
