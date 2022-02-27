package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/modelsout"
	"back-account/src/api/utils/channels"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type repositoryStoreActionCRUD struct {
	db *gorm.DB
}

func NewRepositoryStoreActionCRUD(db *gorm.DB) *repositoryStoreActionCRUD {

	return &repositoryStoreActionCRUD{db}
}

func (r *repositoryStoreActionCRUD) Save(store models.StoreAction) (models.StoreAction, error) {

	var err error
	store.CreatedAt = time.Now()
	done := make(chan bool)
	var code int
	if store.Type==1 {
		r.db.Raw(`select max(code) from store_actions where  company_id=? 
	and year_id=? and type=? `,
	 store.CompanyID, store.YearID, store.Type).Scan(&code)
	} else if store.Type==2 && (store.StoreActionTypeID ==8 || store.StoreActionTypeID ==10 || store.StoreActionTypeID ==13)  {

	
	r.db.Raw(`select max(code) from store_actions where  company_id=? 
	and year_id=? and type=? and store_action_type_id in (8,10,13) `,
	 store.CompanyID, store.YearID, store.Type).Scan(&code)
	} else  {
		r.db.Raw(`select max(code) from store_actions where  company_id=? 
		and year_id=? and type=? and store_action_type_id in (7,14) `,
		 store.CompanyID, store.YearID, store.Type).Scan(&code)
	}

	if code == 0 {
		code = 1
	} else {
		code = code + 1
	}

	store.Code = code

	go func(ch chan<- bool) {
		err = r.db.Create(&store).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return store, nil
	}

	return models.StoreAction{}, err

}

func (r *repositoryStoreActionCRUD) FindAll(search string, typeaction int, yearid int, companyid int, storeid int) ([]models.StoreAction, error) {

	var err error
	storeis := []models.StoreAction{}
	done := make(chan bool)
	fmt.Println(search)
	go func(ch chan<- bool) {

		if len(search) != 0 {
			err = r.db.Debug().Model(&models.StoreAction{}).Where(search).Where(" year_id=? and company_id=? and type=? and store_id=?", yearid, companyid, typeaction, storeid).Find(&storeis).Error
		} else {
			err = r.db.Debug().Model(&models.StoreAction{}).Where("year_id=? and company_id=? and type=? and store_id=?", yearid, companyid, typeaction, storeid).Find(&storeis).Error
		}

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

func (r *repositoryStoreActionCRUD) GetCountObject(companyid int, yearid int, uid int) (float32, error) {

	var err error
	var count float32
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Raw("call getcountstoreobject(?,?,?)", companyid, yearid, uid).Scan(&count)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return count, nil
	}

	return 0, err

}
func (r *repositoryStoreActionCRUD) GetKardex(companyid int, yearid int, uid int) ([]modelsout.Kradex, error) {

	var err error
	kardex := []modelsout.Kradex{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Raw("call kardex(?,?,?)", companyid, yearid, uid).Scan(&kardex)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return kardex, nil
	}

	return nil, err

}
func (r *repositoryStoreActionCRUD) GetRemObject(companyid int, yearid int, uid int) ([]modelsout.Remobject, error) {

	var err error
	remobjects := []modelsout.Remobject{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Raw("call getremobject(?,?,?)", companyid, yearid, uid).Scan(&remobjects)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return remobjects, nil
	}

	return nil, err

}
func (r *repositoryStoreActionCRUD) GetPriceObject(companyid int, yearid int, objectid int, solardate string) (float64, error) {

	var err error
	var price float64
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Raw("call getpricefromstoreid(?,?,?,?)", companyid, yearid, objectid, solardate).Scan(&price)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return price, nil
	}

	return 0, err

}
func (r *repositoryStoreActionCRUD) GetStoreIdbyDocumentID(documentid int) (int, int, error) {

	var err error
	var storeid int = 0
	var typestoreid int = 0
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Model(models.StoreAction{}).Select("id").Where("document_number=?", documentid).Take(&storeid)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storeid, typestoreid, nil
	}

	return 0, 0, err

}

func (r *repositoryStoreActionCRUD) Getsend(companyid int, yearid int) ([]modelsout.StoreActiondoc, error) {

	var err error

	done := make(chan bool)
	res := []modelsout.StoreActiondoc{}
	go func(ch chan<- bool) {
		result := r.db.Raw("call havaleanbar(?,?)", companyid, yearid).Scan(&res)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return res, nil
	}

	return []modelsout.StoreActiondoc{}, err

}

func (r *repositoryStoreActionCRUD) Getrecive(companyid int, yearid int) ([]modelsout.StoreActiondoc, error) {

	var err error

	done := make(chan bool)
	res := []modelsout.StoreActiondoc{}
	go func(ch chan<- bool) {
		result := r.db.Raw("call residanbar(?,?)", companyid, yearid).Scan(&res)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return res, nil
	}

	return []modelsout.StoreActiondoc{}, err

}
func (r *repositoryStoreActionCRUD) FindById(uid uint32) (models.StoreAction, error) {

	var err error
	store := models.StoreAction{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Model(&models.StoreAction{}).Where("id=?", uid).Preload("StoreActionRow").Find(&store)
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

	return models.StoreAction{}, err

}

func (r *repositoryStoreActionCRUD) GetStoreActionRows(uid uint32) ([]models.StoreActionRow, error) {

	var err error
	storeactionrows := []models.StoreActionRow{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Raw("call getstoreactionrows(?)", uid).Scan(&storeactionrows)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storeactionrows, nil
	}

	return nil, err

}

func (r *repositoryStoreActionCRUD) Update(store models.StoreAction) (int64, error) {

	var err error
	store.CreatedAt = time.Now()

	err = r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&store).Error

	err = r.db.Model(&store).Association("StoreActionRow").Replace(&store.StoreActionRow)
	if err != nil {
		return 0, err
	}
	return 1, err

}

func (r *repositoryStoreActionCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.StoreAction{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryStoreActionCRUD) GetAll(companyid int, yearid int, typestore int) ([]models.StoreAc, error) {

	var err error
	storeactions := []models.StoreAc{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Raw("call getstoreac(?,?,?)", companyid, yearid, typestore).Scan(&storeactions).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.Ok(done) {
		return storeactions, nil
	}
	return []models.StoreAc{}, err
}
