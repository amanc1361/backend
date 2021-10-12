package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryStoreSubGroupCRUD struct {
	db *gorm.DB
}

func NewRepositoryStoreSubGroupCRUD(db *gorm.DB) *repositoryStoreSubGroupCRUD {

	return &repositoryStoreSubGroupCRUD{db}
}

func (r *repositoryStoreSubGroupCRUD) Save(storesubgroup models.StoreSubGroup) (models.StoreSubGroup, error) {

	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Create(&storesubgroup).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storesubgroup, nil
	}

	return models.StoreSubGroup{}, err

}

func (r *repositoryStoreSubGroupCRUD) FindAll() ([]models.StoreSubGroup, error) {

	var err error
	storesubgroupis := []models.StoreSubGroup{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Find(&storesubgroupis)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storesubgroupis, nil
	}

	return nil, err

}

func (r *repositoryStoreSubGroupCRUD) FindById(uid uint32) (models.StoreSubGroup, error) {

	var err error
	storesubgroup := models.StoreSubGroup{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&storesubgroup)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storesubgroup, nil
	}

	return models.StoreSubGroup{}, err

}

func (r *repositoryStoreSubGroupCRUD) Update(storesubgroup models.StoreSubGroup) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Save(&storesubgroup)

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryStoreSubGroupCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.StoreSubGroup{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}
