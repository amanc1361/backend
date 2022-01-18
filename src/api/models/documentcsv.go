package models

type DocumnetCsv struct {
	DocumentNumber int
	SolarDate      string
	LedgerCode     string
	LedgerName     string
	SubLedgerCode  string
	SubLedgerName  string
	DetailedCode   string
	DetailedName   string
	Debtor         int64
	Creditor       int64
}
