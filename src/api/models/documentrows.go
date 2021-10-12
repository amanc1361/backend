package models

import "gorm.io/gorm"

type DocumentRows struct {
	gorm.Model

	DocumentID    uint
	Description   string
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
	Detailed1ID   uint64
	Detailed1Code uint64
	Detailed1Name string
	Detailed2ID   uint64
	Detailed2Code uint64
	Detailed2Name string
	Detailed3ID   uint64
	Detailed3Code uint64
	Detailed3Name string
}
