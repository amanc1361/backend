package models

import "gorm.io/gorm"

type Bilan struct {
	gorm.Model
	Name          string `json:"name"`
	Code          int64  `json:"code"`
	Firstdebtor   int64  `json:"first_debtor"`
	FirstCreditor int64  `json:"first_creditor"`
	PreDebtor     int64  `json:"pre_debtor"`
	PreCreditor   int64  `json:"pre_creditor"`
	Debtor        int64  `json:"debtor"`
	Creditor      int64  `json:"creditor"`
	Remdebtor     int64
	Remcreditor   int64
	Modeltype     int
}
