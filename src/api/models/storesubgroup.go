package models

import "time"

type StoreSubGroup struct {
	ID           uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Name         string     `gorm:"size:100;not null;" json:"name"`
	Code         int        `gorm:"not null;" json:"code"`
	CompanyID    int        `gorm:"not null" json:"company_id"`
	StoreGroupID int        `gorm:"not null" json:"store_group_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

type PublicStoreSubGroup struct {
	ID           uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name         string `gorm:"size:100;not null;" json:"name"`
	Code         int    `gorm:"not null;" json:"code"`
	StoreGroupID int    `gorm:"not null;" json:"store_group_id"`
}

type StoreSubGroups []StoreSubGroup

func (storesubgroups StoreSubGroups) PublicStoreSubGroups() []interface{} {
	result := make([]interface{}, len(storesubgroups))
	for index, Storesubgroup := range storesubgroups {
		result[index] = Storesubgroup.PublicStoreSubGroup()
	}
	return result
}

func (u *StoreSubGroup) PublicStoreSubGroup() interface{} {
	return &PublicStoreSubGroup{
		ID:           u.ID,
		Name:         u.Name,
		Code:         u.Code,
		StoreGroupID: u.StoreGroupID,
	}
}
