package models

import (
	// "back-account/src/api/security"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID       int32  `gorm:"primary_key;auto_increment" json:"id"`
	UserName string `gorm:"size:20;not null;unique" json:"username"`
	Password string `gorm:"size:120;not null;unique" json:"password"`
	Name     string `gorm:"size:60;not null" json:"name"`
	Family   string `gorm:"size:80;not null" json:"family"`
	RoleId   int32  `gorm:"not null" json:"role_id"`
	// CreatedAt time.Time `gorm:"defult:current_timestamp()" json:"created_at"`
	// UpdatedAt time.Time `gorm:"defult:current_timestamp()" json:"updated_at"`
}

// func (u *User) BeforeSave(*gorm.DB) (err error) {

// 	hashpassword, err := security.Hash(u.Password)
// 	if err != nil {
// 		return err
// 	}

// 	u.Password = string(hashpassword)
// 	return

// }
