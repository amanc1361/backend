package models

import (
	"gorm.io/gorm"
)

// StoreActionRow Save in database
type StoreActionRow struct {
	gorm.Model

	StoreActionID   uint
	StoreObjectID   int     `json:"store_object_id"`
	StoreObjectName string  `json:"store_object_name"`
	StoreID         int     `json:"store_id"`
	Subcount        int     `json:"subcount"`
	Countin         float32 `gorm:"defalut:0" json:"countin"`
	Countout        float32 `gorm:"defalut:0" json:"countout"`
	UnitID          int     `json:"unit_id"`
	Price           int     `json:"price"`
	UnitName        string  `json:"unit_name"`
	StoreObjectCode int     `json:"store_object_code"`

}

// PublicStoreActionRow for response to client
type PublicStoreActionRow struct {
	StoreActionID   uint
	StoreObjectID   int     `json:"store_object_id"`
	StoreObjectName string  `json:"store_object_name"`
	StoreID         int     `json:"store_id"`
	Subcount        int     `json:"subcount"`
	Countin         float32 `json:"countin"`
	Countout        float32 `json:"countout"`
	UnitID          int     `json:"unit_id"`
	Price           int     `json:"price"`
	UnitName        string  `json:"unit_name"`
    Tax float32 `json:"tax"`
	ValueAdd float32  `json:"value_add"`
	StoreObjectCode int     `json:"store_object_code"`

}

// StoreActionRows list of StoreActionRow
type StoreActionRows []StoreActionRow

// PublicStoreActionRow get PublicStoreActionRow
func (u *StoreActionRow) PublicStoreActionRow() interface{} {
	return &StoreActionRow{

		StoreActionID:   u.StoreActionID,
		StoreObjectID:   u.StoreObjectID,
		StoreObjectName: u.StoreObjectName,
		Subcount:        u.Subcount,
		Countin:         u.Countin,
		Countout:        u.Countout,
		UnitID:          u.UnitID,
		Price:           u.Price,
		UnitName:        u.UnitName,
		StoreID:         u.StoreID,
	 
		StoreObjectCode: u.StoreObjectCode,

	}
}

func (storeactionrows StoreActionRows) PublicStoreActionRows() []interface{} {
	result := make([]interface{}, len(storeactionrows))
	for index, storeactionrow := range storeactionrows {
		result[index] = storeactionrow
	}
	return result

}
