package models

import (
	"gorm.io/gorm"
)

// StoreActionType the create object to Save StoreActionType to database
type StoreActionType struct {
	gorm.Model
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:100;not null;" json:"name"`
	Type      int    `json:"type"`
	CompanyID int    `gorm:"not null" json:"company_id"`
}

// PublicStoreActionType the create interface for response from database
type PublicStoreActionType struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:100;not null;" json:"name"`
	Type      int    `json:"type"`
	CompanyID int    `gorm:"not null" json:"company_id"`
}

// StoreActionTypeis List of StoreActionType
type StoreActionTypeis []StoreActionType

// PublicStoreActionTypeis Get All Stroe
func (stories StoreActionTypeis) PublicStoreActionTypeis() []interface{} {
	result := make([]interface{}, len(stories))
	for index, store := range stories {
		result[index] = store
	}
	return result
}

// PublicStoreActionType Get StoreActionType
func (u *StoreActionType) PublicStoreActionType() interface{} {
	return &StoreActionType{
		ID:        u.ID,
		Name:      u.Name,
		Type:      u.Type,
		CompanyID: u.CompanyID,
	}
}
