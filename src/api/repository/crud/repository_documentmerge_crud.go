package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"
	"time"

	"gorm.io/gorm"
)

type DocumentMergeRepository struct {
	db *gorm.DB
}

func NewDocumentMergeRepository(db *gorm.DB) *DocumentMergeRepository {
	return &DocumentMergeRepository{db}
}

func (r *DocumentMergeRepository) Save(documentmerge models.DocumentMerge) (models.DocumentMerge, error) {

	var err error

	done := make(chan bool)
	var _solardate string = ""
	var _docnumner int = 0

	documentmerge.CreatedAt = time.Now()
	switch documentmerge.DocumentTypeID {

	case 2:
		if len(documentmerge.SolarFrom) == 10 && len(documentmerge.SolarTo) == 10 {
			r.db.Model(models.Document{}).Where("solar_date>=? and solar_date<=? and company_id=? and year_id=?",
				documentmerge.SolarFrom, documentmerge.SolarTo, documentmerge.CompanyID, documentmerge.YearID).Select("max(solar_date)").Scan(_solardate)

		} else if documentmerge.DocFrom != 0 && documentmerge.DocTo != 0 {
			r.db.Model(models.Document{}).Where("document_number>=? and document_number<=? and company_id=? and year_id=?",
				documentmerge.DocFrom, documentmerge.DocTo, documentmerge.CompanyID, documentmerge.YearID).Select("max(solar_date)").Scan(&_solardate)

		}

	}
	if len(_solardate) != 0 {
		documentmerge.SolarDate = _solardate
	}

	err = r.db.Model(models.DocumentMerge{}).Where("company_id=? and year_id=?", documentmerge.CompanyID, documentmerge.YearID).Select("max(document_number)").Scan(&_docnumner).Error
	// if err != nil {
	// 	return models.DocumentMerge{}, err
	// }
	documentmerge.DocumentNumber = uint(_docnumner) + 1
	go func(ch chan<- bool) {
		err = r.db.Save(&documentmerge).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.Ok(done) {
		return documentmerge, nil
	}
	return models.DocumentMerge{}, err

}

func (r *DocumentMergeRepository) FindAll(companyid int, yearid int) ([]models.DocumentMerge, error) {
	var err error
	documentmerges := []models.DocumentMerge{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Model(models.DocumentMerge{}).Where("company_id=? and year_id=?", companyid, yearid).Find(&documentmerges).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.Ok(done) {
		return documentmerges, nil
	}
	return []models.DocumentMerge{}, err
}

func (r *DocumentMergeRepository) Delete(id int) (int64, error) {

	done := make(chan bool)
	var res *gorm.DB
	go func(ch chan<- bool) {
		res = r.db.Where("id=?", id).Delete(&models.DocumentMerge{})
		ch <- true
	}(done)
	if channels.Ok(done) {
		if res.Error != nil {
			return 0, res.Error
		}
	}
	return res.RowsAffected, nil
}

func (r *DocumentMergeRepository) GetLastDocumetMergeCode(companyid int, yearid int) (int, error) {
	var err error
	var code int
	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Model(models.DocumentMerge{}).Where("comany_id=? and year_id=?", companyid, yearid).Select(" max(DocumentNumber)").Scan(&code).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.Ok(done) {
		return code, nil
	}
	return 0, err
}

func (r *DocumentMergeRepository) GetLedgerDocument(docmerge models.DocumentMerge) ([]models.LedgerDocument, error) {
	var err error
	done := make(chan bool)
	var treport int = 1
	if len(docmerge.SolarFrom) > 8 && len(docmerge.SolarTo) > 8 {
		treport = 2
	}

	ledgerdocuments := []models.LedgerDocument{}
	go func(ch chan<- bool) {
		err = r.db.Debug().Raw("call getledgerdocument(?,?,?,?,?,?,?)", docmerge.CompanyID,
			docmerge.YearID,
			treport,
			docmerge.DocFrom,
			docmerge.DocTo,
			docmerge.SolarFrom,
			docmerge.SolarTo,
		).Scan(&ledgerdocuments).Error
		if err != nil {

			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.Ok(done) {
		return ledgerdocuments, nil
	}

	return nil, err
}
