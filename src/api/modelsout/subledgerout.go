package modelsout

type SubledgerOut struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Code          uint   `json:"code"`
	OwnerDetailed bool   `json:"owner_detailed"`
	LedgerID      uint   `json:"ledger_id"`
	Nature        byte   `json:"nature"`
	CompanyID     int    `json:"company_id"`
}
