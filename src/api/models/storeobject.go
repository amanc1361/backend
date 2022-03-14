package models

import (
	"gorm.io/gorm"
)

// StoreObject  this interface for save store object in database
type StoreObject struct {
	gorm.Model
	ID              uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name            string `gorm:"size:100;not null;" json:"name"`
	Code            int    `gorm:"not null;" json:"code"`
	StoreID         int    `gorm:"not null;" json:"store_id"`
	CompanyID       int    `gorm:"not null" json:"company_id"`
	GroupID         int    `json:"group_id"`
	SubGroupID      int    `json:"sub_group_id"`
	ChildSubGroupID int    `json:"child_sub_group_id"`
	UnitID          int    `json:"unit_id"`
	Mincount        int    `json:"mincount"`
	Maxcount        int    `json:"maxcount"`
	Tax float32 `json:"tax"`
	ValueAdd float32 `json:"value_add"`
	PriceType       int    `json:"price_type"`
	DiscountType    int    `json:"discount_type"`
	Capacity  float32 `json:"capacity"`

}

// PublicStoreObject this interface for response from database to client
type PublicStoreObject struct {
	ID              uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name            string `gorm:"size:100;not null;" json:"name"`
	Code            int    `gorm:"not null;" json:"code"`
	StoreID         int    `gorm:"not null;" json:"store_id"`
	StoreName       string `json:"store_name"`
	CompanyID       int    `gorm:"not null" json:"company_id"`
	GroupID         int    `json:"group_id"`
	StoreGroupName  string `json:"store_group_name"`
	SubGroupID      int    `json:"sub_group_id"`
	ChildSubGroupID int    `json:"child_sub_group_id"`
	UnitID          int    `json:"unit_id"`
	UnitName        string `json:"unit_name"`
	Mincount        int    `json:"mincount"`
	Maxcount        int    `json:"maxcount"`
	PriceType       int    `json:"price_type"`
	DiscountType    int    `json:"discount_type"`
	Tax float32 `json:"tax"`
	ValueAdd float32 `json:"value_add"`
	Capacity  float32 `json:"capacity"`
}

// StoreObjects list of StoreObject
type StoreObjects []StoreObject

func (storeobjects StoreObjects) PublicStoreObjects() []interface{} {
	result := make([]interface{}, len(storeobjects))
	for index, storeobject := range storeobjects {
		result[index] = storeobject
	}
	return result

}

//PublicStoreObject send custom data to client
func (u *StoreObject) PublicStoreObject() interface{} {
	return &PublicStoreObject{
		ID:              u.ID,
		Name:            u.Name,
		Code:            u.Code,
		StoreID:         u.StoreID,
		CompanyID:       u.CompanyID,
		GroupID:         u.GroupID,
		SubGroupID:      u.SubGroupID,
		ChildSubGroupID: u.ChildSubGroupID,
		UnitID:          u.UnitID,
		Mincount:        u.Mincount,
		Maxcount:        u.Maxcount,
		PriceType:       u.PriceType,
		DiscountType:    u.DiscountType,
		Capacity: u.Capacity,
	}
}
