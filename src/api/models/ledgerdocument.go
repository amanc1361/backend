package models

type LedgerDocument struct {
	Code     int    `json:"code"`
	Name     string `json:"name"`
	Debtor   int    `json:"debtor"`
	Creditor int    `json:"creditor"`
}
