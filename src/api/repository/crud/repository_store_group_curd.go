package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryStoreGroupCRUD struct {
	db *gorm.DB
}

func NewRepositoryStoreGroupCRUD(db *gorm.DB) *repositoryStoreGroupCRUD {

	return &repositoryStoreGroupCRUD{db}
}

func (r *repositoryStoreGroupCRUD) Save(storegroup models.StoreGroup) (models.StoreGroup, error) {

	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Create(&storegroup).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storegroup, nil
	}

	return models.StoreGroup{}, err

}

func (r *repositoryStoreGroupCRUD) FindAll(companyid int, search string) ([]models.StoreGroup, error) {

	var err error
	storegroupis := []models.StoreGroup{}
	done := make(chan bool)
	var result *gorm.DB
	go func(ch chan<- bool) {
		if len(search) == 0 {
			result = r.db.Model(&models.StoreGroup{}).Where("company_id=?", companyid).Find(&storegroupis)
		} else {
			result = r.db.Model(&models.StoreGroup{}).Where("company_id=? and (name LIKE ? or  convert(code,char) collate utf8mb4_persian_ci like ?)", companyid, "%"+search+"%", "%"+search+"%").Find(&storegroupis)
		}
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storegroupis, nil
	}

	return nil, err

}

func (r *repositoryStoreGroupCRUD) FindById(uid uint32) (models.StoreGroup, error) {

	var err error
	storegroup := models.StoreGroup{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&storegroup)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storegroup, nil
	}

	return models.StoreGroup{}, err

}

func (r *repositoryStoreGroupCRUD) Update(storegroup models.StoreGroup) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Save(&storegroup)

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryStoreGroupCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.StoreGroup{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}
