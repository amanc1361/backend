package models

import "time"

// Unit  واحد اندازه گیری برای کالاهای انبار
type Unit struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string     `gorm:"size:100;not null;" json:"name"`
	SubName   string     `gorm:"size:100" json:"sub_name"`
	NameCount int        `json:"name_count"`
	CompanyID int        `gorm:"not null" json:"company_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// PublicUnit روجی برای واحد اندازه گیری
type PublicUnit struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:100;not null;" json:"name"`
	SubName   string `gorm:"size:100" json:"sub_name"`
	NameCount int    ` json:"name_count"`
	CompanyID int    `gorm:"not null" json:"company_id"`
}

// Units  یک لیست از واحدهای اندازه گیری
type Units []Unit

// PublicUnits  تابی برای برگرداندن لیست واحدها
func (units Units) PublicUnits() []interface{} {
	result := make([]interface{}, len(units))
	for index, unit := range units {
		result[index] = unit.PublicUnit()
	}
	return result

}

// PublicUnit یک تابع برای نمایش دلخواه از خروجی دیتا بیس برای نمایش واحد اندازه گیری
func (u *Unit) PublicUnit() interface{} {
	return &PublicUnit{
		ID:        u.ID,
		Name:      u.Name,
		SubName:   u.SubName,
		NameCount: u.NameCount,
		CompanyID: u.CompanyID,
	}
}
