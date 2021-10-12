package auto

import (
	"back-account/src/api/database"
	"back-account/src/api/models"
	"log"
)

func Load() {
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		log.Fatal(err)
	}

	// err = db.Migrator().DropTable(&models.User{}, &models.Year{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = db.AutoMigrate(&models.User{}, &models.Year{}, &models.CompanyType{}, &models.Group{}, &models.Ledger{},
		&models.Document{}, &models.DocumentRows{}, &models.SubLedger{}, &models.Detailed{}, &models.DocumentType{},
		&models.Company{}, &models.Group{}, &models.Ledger{},
		&models.Store{}, &models.StoreGroup{}, &models.StoreSubGroup{}, &models.StoreChildSubGroup{},
		&models.StoreObject{},
		&models.StoreActionType{},
		&models.StorePerson{},
		&models.StorePersonRecive{},
		&models.StoreAction{},
		&models.StoreActionRow{},
		&models.Unit{}, &models.Person{}, &models.UesrRelatedCompany{},
	)

	if err != nil {
		log.Fatal(err)
	}

	// for _, user := range users {
	// 	err = db.Create(&user).Error
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// for _, year := range years {
	// 	err = db.Create(&year).Error
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
}

// for _, companies := range company {
// 	err = db.Create(&companies).Error
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	console.Preety(companies)
// }

// for _, companyType := range companytypes {
// 	err = db.Create(&companyType).Error
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	console.Preety(companyType)
// }
