package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryWarehousingCRUD struct {
	db *gorm.DB
}

func NewRepositoryWareHousingCRUD(db *gorm.DB) *repositoryWarehousingCRUD {
	return &repositoryWarehousingCRUD{db}
}

func (r *repositoryWarehousingCRUD) Save(warehousing models.WareHousing) (models.WareHousing, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Create(&warehousing).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.Ok(done) {
		return warehousing, nil
	}
	return models.WareHousing{}, err
}

func (r *repositoryWarehousingCRUD) FindAll(comapnyid int, yearid int) ([]models.WareHousing, error) {
	var err error
	done := make(chan bool)
	warehousings := []models.WareHousing{}
	go func(ch chan<- bool) {
		err = r.db.Where("company_id=? and year_id=?", comapnyid, yearid).Find(&warehousings).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)
	if channels.Ok(done) {
		return warehousings, nil
	}
	return []models.WareHousing{}, err

}
func (r *repositoryWarehousingCRUD) Update(warehouing models.WareHousing) (int64, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&warehouing).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.Ok(done) {
		return 1, nil
	}
	return 0, err
}

func (r *repositoryWarehousingCRUD) GetWareHousingById(id int ) (models.WareHousing, error) {
	var err error
	done := make(chan bool)
	warehousing := models.WareHousing{}
	go func(ch chan<- bool) {
		err = r.db.Where("id=?", id).Take(&warehousing).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.Ok(done) {
		return warehousing, nil
	}
	return models.WareHousing{}, err
}