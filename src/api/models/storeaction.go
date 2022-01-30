package models

import (
	"gorm.io/gorm"
)

// StoreAction Save Action Receve and send object from store
type StoreAction struct {
	gorm.Model

	Description       string `gorm:"size:200;not null;" json:"description"`
	Code              int    `gorm:"not null" json:"code"`
	CompanyID         int    `gorm:"not null" json:"company_id"`
	Type              int    `gorm:"not null" json:"type"`
	SolarDate         string `json:"solar_date"`
	StoreID           int    `json:"store_id"`
	DocumentNumber    int    `json:"document_number"`
	InvoiceID         int    `json:"invoice_id"`
	UserID            int    `json:"user_id"`
	StoreActionTypeID int    `json:"store_action_type_id"`
	StorePersonID     int    `json:"store_person_id"`
	YearID            int    `gorm:"not null" json:"year_id"`
	Tax				  int    `json:"tax"`
	Discount int `json:"discount"`
	CostCenterId int `json:"cost_center_id"`
	CostCenterName string `json:"cost_center_name"`
	StoreActionRow    []StoreActionRow
}

// PublicStoreAction custom response from data
type PublicStoreAction struct {
	Description       string `gorm:"size:200;not null;" json:"description"`
	Code              int    `gorm:"not null;" json:"code"`
	CompanyID         int    `gorm:"not null" json:"company_id"`
	Type              int    `json:"type"`
	DocumentID        int    `json:"document_id"`
	StoreID           int    `json:"store_id"`
	InvoiceID         int    `json:"invoice_id"`
	SolarDate         string `json:"solar_date"`
	UserID            int    `json:"user_id"`
	YearID            int    `json:"year_id"`
	StoreActionTypeID int    `json:"store_action_type_id"`
	StorePersonID     int    `json:"store_person_id"`
	Tax int `json:"tax"`
	Discount int `json:"discount"`
	CostCenterId int `json:"cost_center_id"`
	CostCenterName string `json:"cost_center_name"`
	StoreActionRow    []StoreActionRow
}

// StoreActions List of StoreAction
type StoreActions []StoreAction

// PublicStoreAction get StroeAction
func (u *StoreAction) PublicStoreAction() interface{} {
	return &StoreAction{
		Description:       u.Description,
		Code:              u.Code,
		CompanyID:         u.CompanyID,
		Type:              u.Type,
		SolarDate:         u.SolarDate,
		UserID:            u.UserID,
		YearID:            u.YearID,
		Tax: u.Tax,
		Discount: u.Discount,
		CostCenterId: u.CostCenterId,
		CostCenterName: u.CostCenterName,
		StoreActionTypeID: u.StoreActionTypeID,
		StoreActionRow:    u.StoreActionRow,
	}

}

// PublicStoreActions get all StoreActions
func (storeactions StoreActions) PublicStoreActions() []interface{} {
	result := make([]interface{}, len(storeactions))
	for index, storeaction := range storeactions {
		result[index] = storeaction
	}
	return result
}
