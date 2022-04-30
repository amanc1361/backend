package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/modelsout"
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

func(r *repositoryStoreCRUD) GetStories(companyid int,yearid int)([]models.StoriesRem,error) {
	var err error
	 done:=make(chan bool)
	 stories:=[]models.StoriesRem{}

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
	 return []models.StoriesRem{},err

}

func(r *repositoryStoreCRUD) GetStoreWithObject(companyid int,yearid int,storeid int,reportrtpe int)([]models.StoreRemObjects,error) {
	var err error
	 done:=make(chan bool)
	 stories:=[]models.StoreRemObjects{}

	 go func (ch chan<-bool)  {
		 err=r.db.Raw("call getobjectsbystoreid(?,?,?,?)",companyid,yearid,storeid,reportrtpe).Take(&stories).Error
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

func (r *repositoryStoreCRUD) GetRemObjectByStoreId(companyid int,yearid int,storeid int)([]modelsout.Remobject,error) {
	var err error 
	done:=make(chan bool) 
	remObjects:=[]modelsout.Remobject{}
	go func (ch chan<-bool) {
		err=r.db.Table("store_actions").
		Select("store_objects.code,store_objects.name, sum(countin)-sum(countout) as rem").
		Joins("store_action_rows on store_actions.id=store_action_rows.store_action_id").
		Joins("store_objects on store_action_rows.store_object_id=store_objects.id").
		Where("store_actions.company_id=? and store_actions.year_id=? and store_objects.store_id=?",companyid,yearid,storeid).
		Group("store_objects.code,store_objects.name").
		Habing("rem>0").Order("store_objects.code").Find(&remObjects).Error
		if err!=nil {
			ch<-false
			return
		}
		ch<-true
		return
		
	}
	if channels.Ok(done) {
		return remObjects,nil
	}
	return []modelsout.remObjects{},err
}
