package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"
	"errors"
	"time"

	"gorm.io/gorm"
)

type repositoryLedgerCRUD struct {
	db *gorm.DB
}

func NewRepositoryLedgerCRUD(db *gorm.DB) *repositoryLedgerCRUD {

	return &repositoryLedgerCRUD{db}
}

func (r *repositoryLedgerCRUD) Save(ledger models.Ledger) (models.Custledger, error) {

	var err error

	done := make(chan bool)
	res := models.Custledger{}
	go func(ch chan<- bool) {
		err = r.db.Create(&ledger).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		r.db.Raw("call ledgersbyid(?)", ledger.ID).Take(&res)

		return res, nil
	}

	return models.Custledger{}, err

}

func (r *repositoryLedgerCRUD) FindAll(companyid int, search string) ([]models.Custledger, error) {

	var err error
	ledgers := []models.Custledger{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Raw("call getledgers(?,?)", companyid, search).Scan(&ledgers)

		err = result.Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)
	if channels.Ok(done) {

		return ledgers, nil
	}
	return nil, err
}
func (r *repositoryLedgerCRUD) FindById(uid uint32) (models.Custledger, error) {
	var err error
	ledger := models.Custledger{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		result := r.db.Raw("call ledgersbyid(?)", uid).Take(&ledger)
		err = result.Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.Ok(done) {
		return ledger, nil
	}
	return models.Custledger{}, err
}
func (r *repositoryLedgerCRUD) Update(ledger models.Ledger) (int64, error) {
	var result *gorm.DB
	done := make(chan bool)
	ledger.CreatedAt = time.Now()
	go func(ch chan<- bool) {
		result = r.db.Save(&ledger)
		ch <- true
	}(done)
	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}
func (r *repositoryLedgerCRUD) Delete(uid int32) (int64, error) {
	var result *gorm.DB
	done := make(chan bool)
	var count int64 = 0
	go func(ch chan<- bool) {
		result = r.db.Model(&models.DocumentRows{}).Where("ledger_id=?", uid).Count(&count)
		if count == 0 {
			result = r.db.Where("id=?", uid).Delete(&models.Ledger{})
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

func (r *repositoryLedgerCRUD) FindByCode(code uint, companyid int) (models.Custledger, error) {

	var err error
	ledger := models.Custledger{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Raw("call ledgerbycode(?,?)", code, companyid).Take(&ledger)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return ledger, nil
	}

	return models.Custledger{}, err

}

func (r *repositoryLedgerCRUD) GetLastCode(id uint, companyid int) (uint, error) {
	var code uint

	result := r.db.Raw("call getlastledgercode(?,?)", id, companyid).Scan(&code)

	if result.Error != nil {

		return 0, result.Error
	}
	return code, nil
}

func (r *repositoryLedgerCRUD) GetByGroupID(uid uint) ([]models.Custledger, error) {
	var err error
	ledger := []models.Custledger{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		result := r.db.Model(&models.Ledger{}).Where("group_id=?", uid).Find(&ledger)
		err = result.Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.Ok(done) {
		return ledger, nil
	}
	return nil, err
}
