package crud

type repositoryWarehousingCRUD struct {
	db *gorm.DB
}

func NewRepositoryWareHousingCRUD(db *gorm.DB) *repositoryWarehousingCRUD {
	return repositoryWarehousingCRUD{db}
}

func (r *repositoryWarehousingCRUD) Save(models.WareHousing warehousing) (models.WareHousing,error) {
	var err Error
	done:=make(chan bool)
	 go func (ch chan<-bool) {
		 err=r.db.Create(&warehousing).Error
		 if err!=nil {
			 ch<-false
			 return 
		 }
		 ch<-true
	 }(done)
	 if channels.Ok(done) {
		 return warehousing,nil
	 }
	 return models.WareHousing{},err

}

func (r *repositoryWarehousingCRUD) FindAll(comapnyid int, yearid int)([]models.WareHousing,error) {
  var err error 

  done:=make(chan bool)
  warehousings:=[]models.warehousing{}
  go func (ch chan<-bool) {
	  err=r.db.Where("company_id=? and year_id=?",comapnyid,yearid).Find(&warehousings).Error
	  if err!=nil {
		  ch<-false
		  return
	  }
	  ch<-true 

  }(done)
   if channels.Ok(done) {
	   return warehousings,nil
   }
   return []models.WareHousing{},err
	
 }