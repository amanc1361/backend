package models

import "time"

type StoreGroup struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string     `gorm:"size:100;not null;" json:"name"`
	Code      int        `gorm:"not null;" json:"code"`
	CompanyID int        `gorm:"not null" json:"company_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type PublicStoreGroup struct {
	ID   uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:100;not null;" json:"name"`
	Code int    `gorm:"not null;" json:"code"`
}

type StoreGroups []StoreGroup

func (storegroups StoreGroups) PublicStoreGroups() []interface{} {
	result := make([]interface{}, len(storegroups))
	for index, Storegroup := range storegroups {
		result[index] = Storegroup.PublicStoreGroup()
	}
	return result
}

func (u *StoreGroup) PublicStoreGroup() interface{} {
	return &PublicStoreGroup{
		ID:   u.ID,
		Name: u.Name,
		Code: u.Code,
	}
}
