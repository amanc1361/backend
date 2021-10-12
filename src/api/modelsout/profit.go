package modelsout

type Profit struct {
	Debtor        uint64
	Creditor      uint64
	LedgerID      uint64
	LedgerCode    uint64
	LedgerName    string
	SubLedgerID   uint64
	SubLedgerCode uint64
	SubLedgerName string
	DetailedID    uint64
	DetailedCode  uint64
	DetailedName  string
	Nature        int
}
