package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/modelsout"
	"back-account/src/api/utils/channels"
	"errors"
	"time"

	"gorm.io/gorm"
)

type repositorySubLedgerCRUD struct {
	db *gorm.DB
}

func NewRepositorySubLedgerCRUD(db *gorm.DB) *repositorySubLedgerCRUD {

	return &repositorySubLedgerCRUD{db}
}

func (r *repositorySubLedgerCRUD) Save(subledger models.SubLedger) (modelsout.SubledgerOut, error) {

	var err error
	ledgerout := modelsout.SubledgerOut{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Create(&subledger).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		r.db.Raw("call subledgerbyid(?)", subledger.ID).Take(&ledgerout)
		return ledgerout, nil
	}

	return modelsout.SubledgerOut{}, err

}

func (r *repositorySubLedgerCRUD) FindAll(companyid int) ([]modelsout.SubledgerOut, error) {

	var err error
	subledgers := []modelsout.SubledgerOut{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Raw("call getsubledgers(?);", companyid).Scan(&subledgers)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return subledgers, nil
	}

	return nil, err

}

func (r *repositorySubLedgerCRUD) FindById(uid uint32) (modelsout.SubledgerOut, error) {

	var err error
	subledger := modelsout.SubledgerOut{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Raw("call subledgerbyid(?)", uid).Take(&subledger)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return subledger, nil
	}

	return modelsout.SubledgerOut{}, err

}

func (r *repositorySubLedgerCRUD) Update(subledger models.SubLedger) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	subledger.CreatedAt = time.Now()
	go func(ch chan<- bool) {

		result = r.db.Save(&subledger)

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositorySubLedgerCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	var count int64 = 0
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Model(&models.DocumentRows{}).Where("sub_ledger_id=?", uid).Count(&count)
		if count == 0 {
			result = r.db.Where("id=?", uid).Delete(&models.SubLedger{})
		}

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}
	if count > 0 {
		return 0, errors.New("سرفصل انتخابی به دلیل داشتن گردش مالی قابل حذف نیست.")
	}
	return result.RowsAffected, nil

}

func (r *repositorySubLedgerCRUD) FindByCode(code uint, companyid int) (modelsout.SubledgerOut, error) {

	var err error
	subledger := modelsout.SubledgerOut{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Raw("call subledgerbycode(?,?)", code, companyid).Take(&subledger)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return subledger, nil
	}

	return modelsout.SubledgerOut{}, err

}

func (r *repositorySubLedgerCRUD) GetLastCode(codesubledger uint, companyid int) (uint, error) {
	var code uint

	result := r.db.Raw("call getlastsubledgercode(?,?)", codesubledger, companyid).Scan(&code)

	if result.Error != nil {

		return 0, result.Error
	}
	return code, nil
}

func (r *repositorySubLedgerCRUD) GetByLedgerID(uid uint) ([]modelsout.SubledgerOut, error) {

	var err error
	subledgers := []modelsout.SubledgerOut{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Model(&models.SubLedger{}).Where("ledger_id=?", uid).Find(&subledgers)
	
		err = result.Error
         
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return subledgers, nil
	}

	return nil, err

}
