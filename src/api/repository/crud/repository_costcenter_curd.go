package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"
	"errors"

	"gorm.io/gorm"
)
type repositoryCostCenterCrud struct {
	db *gorm.DB
}
func NewRepositoryCostCenter(db *gorm.DB) * repositoryCostCenterCrud {
	return &repositoryCostCenterCrud{db}
}
func (r *repositoryCostCenterCrud) FindAll(companyid int) ([]models.CostCenter,error) {
	var err error
	costcenters:=[]models.CostCenter{}
	done:=make(chan bool)
	go func(ch chan<-bool) 	 {
		 err=r.db.Where("company_id=?",companyid).Find(&costcenters).Error
		 if err!=nil {
			 ch<-false
			 return
		 }
		 ch<-true
 	 }(done)
	 if channels.Ok(done) {
		 return costcenters,nil
	 }
	 return nil,err
}
func (r *repositoryCostCenterCrud) Save(  costcenter models.CostCenter) (models.CostCenter,error) {
	var err error
	done:=make(chan bool) 
	go func (ch chan<-bool)  {
		err=r.db.Create(&costcenter).Error
		if err!=nil {
			ch<-false
			return
		}
	}(done)
	if channels.Ok(done) {
		return costcenter,nil
	}
	return models.CostCenter{},err
}
func (r *repositoryCostCenterCrud) FindByID(id int) (models.CostCenter,error){
	var err error
	costcentr:=models.CostCenter{}
	done:=make(chan bool )
	go func (ch chan<-bool ) {
		err=r.db.Where("id=?",id).First(&costcentr).Error
		if err!=nil {
			ch<-false
			return 
		}
		ch<-true
	}(done)
	if channels.Ok(done) {
		return costcentr,nil
	}
	return models.CostCenter{},err
}
func (r *repositoryCostCenterCrud) Delete( costCenterId int) (int ,error) {
	var err error 
	var count int64
	done:=make(chan bool)

	go func(ch chan<-bool) {
        err=r.db.Model(models.StoreActionRows{}).Where("cost_center_id=?",costCenterId).Count(&count).Error
		if err!=nil {
			ch<-false
			return
		}
		if count!=0 {
			ch<-false
			return 
		}
		if count==0 {
		err=r.db.Where("id=?",costCenterId).Delete(models.CostCenter{}).Error
		if err!=nil {
			ch<-false
			return
		}
		}
		
		ch<-true
	}(done)
	if channels.Ok(done) {
		if count==0 {
		return 1,nil
		} else {
		 return 0,errors.New("این مرکز هزینه قابل حذف نمی باشد")
		}
	}
	return 0,err
}
func (r *repositoryCostCenterCrud) Update(costcenter models.CostCenter)(models.CostCenter,error) {
	var err error
	done:=make(chan bool)
	go func(ch chan<-bool) {
		err=r.db.Model(models.CostCenter{}).Save(&costcenter).Error
		if err!=nil {
			ch<-false
			return
		}
		ch<-true
	}(done) 
	
	if channels.Ok(done) {
		return costcenter,nil
	}
	return models.CostCenter{},err 
}

