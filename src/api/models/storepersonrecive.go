package models

import (
	"gorm.io/gorm"
)

// StorePersonRecive the create object to Save StorePersonRecive to database
type StorePersonRecive struct {
	gorm.Model
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:100;not null;" json:"name"`
	Family    string `gorm:"size:100;not null;" json:"family"`
	Tel       string `gorm:"size:100" json:"tel"`
	Address   string `gorm:"size:200" json:"address"`
	CompanyID int    `gorm:"not null" json:"company_id"`
}

// PublicStorePersonRecive the create interface for response from database
type PublicStorePersonRecive struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:100;not null;" json:"name"`
	Family    string `gorm:"size:100;not null;" json:"family"`
	Tel       string `gorm:"size:100" json:"tel"`
	Address   string `gorm:"size:200" json:"address"`
	CompanyID int    `gorm:"not null" json:"company_id"`
}

// StorePersonReciveis List of StorePersonRecive
type StorePersonReciveis []StorePersonRecive

// PublicStorePersonReciveis Get All Stroe
func (stories StorePersonReciveis) PublicStorePersonReciveis() []interface{} {
	result := make([]interface{}, len(stories))
	for index, store := range stories {
		result[index] = store
	}
	return result
}

// PublicStorePersonRecive Get StorePersonRecive
func (u *StorePersonRecive) PublicStorePersonRecive() interface{} {
	return &StorePersonRecive{
		ID:        u.ID,
		Name:      u.Name,
		Family:    u.Family,
		Tel:       u.Tel,
		Address:   u.Address,
		CompanyID: u.CompanyID,
	}
}
