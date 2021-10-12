package models

type Custledger struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Code      uint   `json:"code"`
	GroupID   uint   `json:"group_id"`
	Nature    byte   `json:"nature"`
	CompanyID int    `json:"company_id"`
}
