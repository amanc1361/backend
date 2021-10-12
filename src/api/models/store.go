package models

import (
	"time"
)

// Store the create object to Save Store to database
type Store struct {
	ID         uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Name       string     `gorm:"size:100;not null;" json:"name"`
	KeeperName string     `gorm:"size:100;not null;" json:"keeper_name"`
	Address    string     `gorm:"size:100;not null;" json:"address"`
	CompanyID  int        `gorm:"not null" json:"company_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

// PublicStore the create interface for response from database
type PublicStore struct {
	ID         uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name       string `gorm:"size:100;not null;" json:"name"`
	CompanyID  int    `gorm:"not null" json:"company_id"`
	KeeperName string `gorm:"size:100;not null;" json:"keeper_name"`
	Address    string `gorm:"size:100;not null;" json:"address"`
}

// Storeis List of Store
type Storeis []Store

// PublicStoreis Get All Stroe
func (stories Storeis) PublicStoreis() []interface{} {
	result := make([]interface{}, len(stories))
	for index, store := range stories {
		result[index] = store
	}
	return result
}

// PublicStore Get Store
func (u *Store) PublicStore() interface{} {
	return &Store{
		ID:         u.ID,
		Name:       u.Name,
		KeeperName: u.KeeperName,
		Address:    u.Address,
		CompanyID:  u.CompanyID,
	}
}
