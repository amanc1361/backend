package models

type GroupTaraz struct {
	GroupCode      int    `json:"group_code"`
	GroupName      string `json:"group_name"`
	Nature         int    `json:"nature"`
	LedgerCode     int    `json:"ledger_code"`
	LedgerName     string `json:"ledger_name"`
	SubLedgerCode  int    `json:"sub_ledger_code"`
	SubLedgerName  string `json:"sub_ledger_name"`
	DetailedCode   int    `json:"detailed_code"`
	DetailedName   string `json:"detailed_name"`
	Debtor         int    `json:"debtor"`
	Creditor       int    `json:"creditor"`
	SolarDate      string `json:"solar_date"`
	DocumentNumber int    `json:"document_number"`
	DocType        int    `json:"doc_type"`
}
