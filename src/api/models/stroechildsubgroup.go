package models

import "time"

type StoreChildSubGroup struct {
	ID              uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Name            string     `gorm:"size:100;not null;" json:"name"`
	Code            int        `gorm:"not null;" json:"code"`
	CompanyID       int        `gorm:"not null" json:"company_id"`
	StoreSubGroupID int        `gorm:"not null" json:"store_sub_group_id"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

type PublicStoreChildSubGroup struct {
	ID              uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name            string `gorm:"size:100;not null;" json:"name"`
	Code            int    `gorm:"not null;" json:"code"`
	StoreSubGroupID int    `gorm:"not null;" json:"store_sub_group_id"`
}

type StoreChildSubGroups []StoreChildSubGroup

func (storechildsubgroups StoreChildSubGroups) PublicStoreChildSubGroups() []interface{} {
	result := make([]interface{}, len(storechildsubgroups))
	for index, Storechildsubgroup := range storechildsubgroups {
		result[index] = Storechildsubgroup.PublicStoreChildSubGroup()
	}
	return result
}

func (u *StoreChildSubGroup) PublicStoreChildSubGroup() interface{} {
	return &PublicStoreChildSubGroup{
		ID:              u.ID,
		Name:            u.Name,
		Code:            u.Code,
		StoreSubGroupID: u.StoreSubGroupID,
	}
}
