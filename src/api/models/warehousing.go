package models

import "gorm.io/gorm"

type WareHousing struct {
	gorm.Model
	StoreId          int                `json:"store_id"`
	SolarDate        string             `json:"solar_date"`
	ReceiptId        int                `json:"receipt_id"`
	TransferId       int                `json:"transfer_id"`
	WareHouseingRows []WareHouseingRows `json:"ware_housing_rows"`
}