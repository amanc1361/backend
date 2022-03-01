package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"
	"time"

	"gorm.io/gorm"
)

type repositoryStoreCRUD struct {
	db *gorm.DB
}

func NewRepositoryStoreCRUD(db *gorm.DB) *repositoryStoreCRUD {

	return &repositoryStoreCRUD{db}
}

func (r *repositoryStoreCRUD) Save(store models.Store) (models.Store, error) {

	var err error

	done := make(chan bool)
	store.CreatedAt = time.Now()
	go func(ch chan<- bool) {
		err = r.db.Save(&store).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return store, nil
	}

	return models.Store{}, err

}

func (r *repositoryStoreCRUD) FindAll(companyid int, search string) ([]models.Store, error) {

	var err error
	storeis := []models.Store{}
	done := make(chan bool)
	var result *gorm.DB

	go func(ch chan<- bool) {
		if len(search) == 0 {
			result = r.db.Model(&models.Store{}).Where("company_id=?", companyid).Find(&storeis)
		} else {
			result = r.db.Model(&models.Store{}).Where("company_id=? and (name LIKE ? or  convert(id,char) collate utf8mb4_persian_ci like ?)", companyid, "%"+search+"%", "%"+search+"%").Find(&storeis)
		}
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storeis, nil
	}

	return nil, err

}

func (r *repositoryStoreCRUD) FindById(uid uint32) (models.Store, error) {

	var err error
	store := models.Store{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&store)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return store, nil
	}

	return models.Store{}, err

}

func (r *repositoryStoreCRUD) Update(store models.Store) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	store.CreatedAt = time.Now()
	go func(ch chan<- bool) {

		result = r.db.Save(&store)

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryStoreCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.Store{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func(r *repositoryStoreCRUD) GetStories(companyid int,yearid int)([]models.Stories,error) {
	var err error
	 done:=make(chan bool)
	 stories:=[]models.Stories{}

	 go func (ch chan<-bool)  {
		 err=r.db.Raw("call getstories(?,?)",companyid,yearid).Take(&stories).Error
		 if err!=nil {
			 ch<-false
			 return
		 }
		 ch<-true
		 
	 }(done)
	 if channels.Ok(done) {
		 return stories,nil
	 }
	 return []models.Stories{},err

}

func(r *repositoryStoreCRUD) GetStoreWithObject(companyid int,yearid int,storeid int)([]models.StoreRemObjects,error) {
	var err error
	 done:=make(chan bool)
	 stories:=[]models.StoreRemObjects{}

	 go func (ch chan<-bool)  {
		 err=r.db.Raw("call getobjectsbystoreid(?,?,?)",companyid,yearid,storeid).Take(&stories).Error
		 if err!=nil {
			 ch<-false
			 return
		 }
		 ch<-true
		 
	 }(done)
	 if channels.Ok(done) {
		 return stories,nil
	 }
	 return []models.StoreRemObjects{},err

}
