package models

type Article struct {
	Document_id     int    `json:"document_id"`
	Document_number int    `json:"document_number"`
	SolarDate       string `json:"solar_date"`
	Description     string `json:"description"`
	Debtor          int64  `json:"debtor"`
	Creditor        int64  `json:"creditor"`
	Nature          int    `json:"nature"`
}
