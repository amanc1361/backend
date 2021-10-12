package modelsout

type Doc struct {
	ID             int
	Debtor         int
	Creditor       int
	DocumentNumber int
	SolarDate      string
	SubLedgerCode  int
	Nature         byte
	Description    string
}
