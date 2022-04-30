package models 
type WareHouseingRows struct {
	gorm.Model
    StoreObjectCode int `json:"store_object_code"`
	StoreObjectName string `json:"store_object_name"`
	Rem int `json:"rem"`
	Conflict int  `json:"conflict"`

	
	
}