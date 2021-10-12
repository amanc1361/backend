package auth

import (
	"back-account/src/api/database"
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"
	"errors"
	"fmt"
)

// SignIn method
func SignIn(username, password string) (string, models.User, error) {
	user := models.User{}
	var err error
	// var db *gorm.DB
	done := make(chan bool)
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	go func(ch chan<- bool) {
		defer close(ch)

		if err != nil {
			ch <- false
			return
		}

		err = db.Debug().Model(models.User{}).Where("user_name = ? and password=?", username, password).First(&user).Error
		if err != nil {
			ch <- false
			return
		}

		if user.Password != password {
			fmt.Println("false compare")
			ch <- false

		}
		// err = security.VerifyPassword(user.Password, password)
		// if err != nil {
		// 	ch <- false
		// 	return
		// }
		ch <- true
	}(done)

	if channels.Ok(done) {
		user.Password = ""
		token, er := GenerateJWT(user)
		if er != nil {
			return "", models.User{}, er
		}
		return token, user, nil
	}

	return "", models.User{}, errors.New("نام کاربری یا رمز عبور اشتباه است")
}
