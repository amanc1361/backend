package models

import (
	"gorm.io/gorm"
)

// StorePerson the create object to Save StorePerson to database
type StorePerson struct {
	gorm.Model
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:100;not null;" json:"name"`
	Family    string `gorm:"size:100;not null;" json:"family"`
	Tel       string `gorm:"size:100" json:"tel"`
	Address   string `gorm:"size:200" json:"address"`
	CompanyID int    `gorm:"not null" json:"company_id"`
}

// PublicStorePerson the create interface for response from database
type PublicStorePerson struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:100;not null;" json:"name"`
	Family    string `gorm:"size:100;not null;" json:"family"`
	Tel       string `gorm:"size:100" json:"tel"`
	Address   string `gorm:"size:200" json:"address"`
	CompanyID int    `gorm:"not null" json:"company_id"`
}

// StorePersonis List of StorePerson
type StorePersonis []StorePerson

// PublicStorePersonis Get All Stroe
func (stories StorePersonis) PublicStorePersonis() []interface{} {
	result := make([]interface{}, len(stories))
	for index, store := range stories {
		result[index] = store
	}
	return result
}

// PublicStorePerson Get StorePerson
func (u *StorePerson) PublicStorePerson() interface{} {
	return &StorePerson{
		ID:        u.ID,
		Name:      u.Name,
		Family:    u.Family,
		Tel:       u.Tel,
		Address:   u.Address,
		CompanyID: u.CompanyID,
	}
}
