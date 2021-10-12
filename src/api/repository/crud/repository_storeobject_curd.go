package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"
	"time"

	"gorm.io/gorm"
)

type repositoryStoreObjectCRUD struct {
	db *gorm.DB
}

func NewRepositoryStoreObjectCRUD(db *gorm.DB) *repositoryStoreObjectCRUD {

	return &repositoryStoreObjectCRUD{db}
}

func (r *repositoryStoreObjectCRUD) Save(storeobjects models.StoreObject) (models.StoreObject, error) {

	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Save(&storeobjects).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storeobjects, nil
	}

	return models.StoreObject{}, err

}

func (r *repositoryStoreObjectCRUD) FindAll() ([]models.StoreObject, error) {

	var err error
	storeobjectsis := []models.StoreObject{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Find(&storeobjectsis)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storeobjectsis, nil
	}

	return nil, err

}
func (r *repositoryStoreObjectCRUD) GetList(companyid int, storeid int) ([]models.PublicStoreObject, error) {

	var err error
	storeobjectsis := []models.PublicStoreObject{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Raw("call getstoreobjects(?,?)", companyid, storeid).Take(&storeobjectsis)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storeobjectsis, nil
	}

	return nil, err

}
func (r *repositoryStoreObjectCRUD) GetCount(companyid int) ([]models.PublicStoreObject, error) {

	var err error
	storeobjectsis := []models.PublicStoreObject{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Raw("call getstoreobjects(?)", companyid).Take(&storeobjectsis)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storeobjectsis, nil
	}

	return nil, err

}

func (r *repositoryStoreObjectCRUD) FindById(uid uint32) (models.StoreObject, error) {

	var err error
	storeobjects := models.StoreObject{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&storeobjects)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storeobjects, nil
	}

	return models.StoreObject{}, err

}

func (r *repositoryStoreObjectCRUD) Update(storeobjects models.StoreObject) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	storeobjects.CreatedAt = time.Now()
	go func(ch chan<- bool) {

		result = r.db.Save(&storeobjects)

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryStoreObjectCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.StoreObject{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}
