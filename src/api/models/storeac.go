package models

type StoreAc struct {
	Id             int    `json:"id"`
	Description    string `json:"description"`
	Name           string `json:"name"`
	PersonName     string `json:"person_name"`
	KeeperName     string `json:"keeper_name"`
	Code           int    `json:"code"`
	CompanyId      int    `json:"company_id"`
	StoreName      string `json:"store_name"`
	SolarDate      string `json:"solar_date"`
	StoreId        int    `json:"store_id"`
	DocumentId     int    `json:"document_id"`
	StorePersionId int    `json:"store_persion_id"`
	YearId         int    `json:"year_id"`
	DocumentNumber int    `json:"document_number"`
}
