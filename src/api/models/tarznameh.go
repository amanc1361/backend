package models

type Tarznameh struct {
	Groupname     string
	Groupcode     string
	Ledgername    string
	Ledgercode    string
	Subledgername string
	Subledgercode string
	Debtor        int64
	Creditor      int64
	Nature        int
}
