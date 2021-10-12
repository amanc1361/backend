package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/modelsout"
	"back-account/src/api/utils/channels"
	"back-account/src/api/utils/paginate"
	"back-account/src/api/utils/pagination"
	"fmt"
	"sync"
	"time"

	"gorm.io/gorm"
)

type documentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) *documentRepository {
	return &documentRepository{db}
}

func (r *documentRepository) Save(document models.Document) (models.Document, error) {
	var err error
	done := make(chan bool)
	var lock sync.Mutex
	lock.Lock()
	go func(ch chan<- bool) {
		document.DocumentNumber, err = r.GetLastCode(document.YearID, document.CompanyID)
		err = r.db.Create(&document).Error
		if err != nil {
			ch <- false
			return
		}
		r.db.Raw("call updatedocumentnumber(?,?)", document.CompanyID, document.YearID)

		ch <- true

	}(done)
	lock.Unlock()
	if channels.Ok(done) {
		return document, nil
	}

	return models.Document{}, err

}

func (r *documentRepository) SortDoc(companyid int, yearid int) (uint, error) {
	var count int
	r.db.Raw("call updatedocumentnumber(?,?)", companyid, yearid).Scan(&count)
	fmt.Println(count)
	return 1, nil

}

func (r *documentRepository) FindAll(search string, yearid uint, companyid int, page int, sort string) (pagination.Pagination, error) {

	var err error
	documents := []models.Document{}

	done := make(chan bool)
	pagination := pagination.Pagination{}
	pagination.Page = page
	pagination.Sort = sort
	var totalrows int64
	go func(ch chan<- bool) {
		defer close(ch)
		if len(search) != 0 {
			r.db.Model(&models.Document{}).Where(search).Where("year_id=? and company_id=?", yearid, companyid).Count(&totalrows)
			r.db.Scopes(paginate.Paginate(totalrows, &pagination, r.db)).Where(search).Where("year_id=? and company_id=?", yearid, companyid).Find(&documents).Order("document_number")
			pagination.Rows = documents
		} else {
			r.db.Model(&models.Document{}).Where("year_id=? and company_id=?", yearid, companyid).Count(&totalrows)
			r.db.Scopes(paginate.Paginate(totalrows, &pagination, r.db)).Where("year_id=? and company_id=?", yearid, companyid).Find(&documents).Order("document_number")
			pagination.Rows = documents

		}

		ch <- true

	}(done)

	if channels.Ok(done) {
		return pagination, nil
	}

	return pagination, err

}
func (r *documentRepository) GetByType(companyid int, yearid int, typedoc int) ([]models.Document, error) {

	var err error
	documents := []models.Document{}

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		r.db.Model(&models.Document{}).Where("year_id=? and company_id=? and document_type_id=?", yearid, companyid, typedoc).Find(&documents)

		ch <- true

	}(done)

	if channels.Ok(done) {
		return documents, nil
	}

	return nil, err

}
func (r *documentRepository) FindByID(uid uint) (models.Document, error) {

	var err error
	documents := models.Document{}

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		r.db.Model(&models.Document{}).Where("id=? ", uid).Take(&documents)
		ch <- true

	}(done)

	if channels.Ok(done) {
		return documents, nil
	}

	return models.Document{}, err

}
func (r *documentRepository) FindRowsByDcoumentid(uid uint) ([]models.DocumentRows, error) {

	var err error
	rows := []models.DocumentRows{}

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		r.db.Raw("call getdocumentrowsbydocumentid(?) ", uid).Scan(&rows)
		ch <- true

	}(done)

	if channels.Ok(done) {
		return rows, nil
	}

	return nil, err

}

func (r *documentRepository) FindByDocumentCode(uid uint, yearid int, companyid int) (models.Document, error) {

	var err error
	documents := models.Document{}

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		t := r.db.Model(&models.Document{}).Where("document_number=? and year_id=? and company_id=?", uid, yearid, companyid).Take(&documents)
		err = t.Error
		if t.Error != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return documents, nil
	}

	return models.Document{}, err

}

func (r *documentRepository) FindByDocumentDescription(des string, yearid int, companyid int) ([]models.Document, error) {

	var err error
	documents := []models.Document{}

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		t := r.db.Model(&models.Document{}).Where("year_id=? and company_id=? and description like %?%", yearid, companyid, des).Preload("DocumentRows").Find(&documents)
		err = t.Error
		if t.Error != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return documents, nil
	}

	return nil, err

}

func (r *documentRepository) Update(uid uint, companyid int, doc models.Document) (uint, error) {

	var err error
	doc.CreatedAt = time.Now()

	err = r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&doc).Error

	err = r.db.Model(&doc).Association("DocumentRows").Replace(&doc.DocumentRows)
	if err != nil {
		return 0, err
	}
	return 1, err

}

func (r *documentRepository) GetLastCode(yearid uint, companyid uint) (uint, error) {
	var code uint

	result := r.db.Raw("call getlastdocumentcode(?,?);", yearid, companyid).Scan(&code)

	if result.Error != nil {

		return 0, result.Error
	}
	return code, nil
}
func (r *documentRepository) GetFirstDate(companyid int, yearid int) (string, error) {
	var solardate string

	result := r.db.Raw("select min(solar_date) from documents where deleted_at is null and company_id=? and year_id=?", companyid, yearid).Take(&solardate)

	if result.Error != nil {

		return "", result.Error
	}
	return solardate, nil
}
func (r *documentRepository) GetLastDate(companyid int, yearid int) (string, error) {
	var solardate string

	result := r.db.Raw("select max(solar_date) from documents where deleted_at is null and company_id=? and year_id=?", companyid, yearid).Take(&solardate)

	if result.Error != nil {

		return "", result.Error
	}
	return solardate, nil
}

func (r *documentRepository) GetInfo(yearid uint, companyid uint) (modelsout.DocInfo, error) {

	docinfo := modelsout.DocInfo{}

	result := r.db.Raw("call getdocinfo(?,?);", yearid, companyid).Scan(&docinfo)

	if result.Error != nil {

		return modelsout.DocInfo{}, result.Error
	}
	return docinfo, nil
}
func (r *documentRepository) Delete(uid uint) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.Document{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}
