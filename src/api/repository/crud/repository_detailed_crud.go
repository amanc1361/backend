package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/modelsout"
	"back-account/src/api/utils/channels"
	"errors"
	"time"

	"gorm.io/gorm"
)

type repositoryDetailedCRUD struct {
	db *gorm.DB
}

func NewRepositoryDetailedCRUD(db *gorm.DB) *repositoryDetailedCRUD {

	return &repositoryDetailedCRUD{db}
}

func (r *repositoryDetailedCRUD) Save(detailed models.Detailed) (models.Detailed, error) {

	var err error

	done := make(chan bool)
	detailed.CreatedAt = time.Now()
	if detailed.Code == 0 {
		detailed.Code, err = r.GetLastCode(int(detailed.CompanyId))
	}
	go func(ch chan<- bool) {
		err = r.db.Save(&detailed).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return detailed, nil
	}

	return models.Detailed{}, err

}

func (r *repositoryDetailedCRUD) FindAll(search string, companyid int) ([]models.Detailed, uint, error) {

	var err error
	var count uint
	var result *gorm.DB
	detaileds := []models.Detailed{}
	done := make(chan bool)

	r.db.Raw("select count(id) from detaileds where company_id=? USE INDEX (PRIMARY);", companyid).Scan(&count)
	go func(ch chan<- bool) {
		if len(search) != 0 {
			result = r.db.Raw("call getdetaileds(?,?)", companyid, search).Scan(&detaileds)
			// result = r.db.Model(&models.Detailed{}).Where("company_id=? and (name LIKE ? or  convert(code,char) collate utf8mb4_persian_ci like ?)", companyid, "%"+search+"%", "%"+search+"%").Find(&detaileds).Order("code")
		} else {
			result = r.db.Model(&models.Detailed{}).Where("company_id=?", companyid).Find(&detaileds).Order("code")
		}

		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return detaileds, count, nil
	}

	return nil, 0, err

}

func (r *repositoryDetailedCRUD) FindById(uid uint32) (models.Detailed, error) {

	var err error
	detailed := models.Detailed{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&detailed)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return detailed, nil
	}

	return models.Detailed{}, err

}

func (r *repositoryDetailedCRUD) Update(uid int32, detailed models.Detailed) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	detailed.CreatedAt = time.Now()
	go func(ch chan<- bool) {

		result = r.db.Model(&models.Detailed{}).Where("id=?", uid).Updates(models.Detailed{Name: detailed.Name,
			IsFloat:     detailed.IsFloat,
			SubLedgerID: detailed.SubLedgerID,
			LedgerID:    detailed.LedgerID, Code: detailed.Code})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryDetailedCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	var count int64
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Debug().Model(&models.DocumentRows{}).Where("detailed_id=?", uid).Count(&count)

		if count == 0 {

			result = r.db.Where("id=?", uid).Delete(&models.Detailed{})

		} else {
			result.Error = errors.New("حساب مورد نظر به دلیل داشتن گردش مالی قابل حذف نیست")
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryDetailedCRUD) FindByCode(uid uint, companyid int) (models.Detailed, error) {

	var err error
	detailed := models.Detailed{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("company_id=? and Code=?", companyid, uid).First(&detailed)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return detailed, nil
	}

	return models.Detailed{}, err

}

func (r *repositoryDetailedCRUD) GetLastCode(companyid int) (uint, error) {
	var code uint

	result := r.db.Raw("call getlastdetailedcode(?)", companyid).Scan(&code)

	if result.Error != nil {

		return 0, result.Error
	}
	return code, nil
}

func (r *repositoryDetailedCRUD) GetFlow(yearid int, detailedid int) ([]modelsout.DetailedFlow, error) {

	detailedflow := []modelsout.DetailedFlow{}

	result := r.db.Raw("call getflowdetaileds(?,?)", yearid, detailedid).Scan(&detailedflow)

	if result.Error != nil {

		return nil, result.Error
	}
	return detailedflow, nil
}
