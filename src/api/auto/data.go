package auto

import "back-account/src/api/models"

var users = []models.User{
	models.User{
		UserName: "admin",
		Password: "Admin8116",
		Name:     "فاطمه",
		Family:   "بهرام میرزایی",
		RoleId:   1,
	},
}

var years = []models.Year{
	models.Year{Name: "سال مالی 1396", CompanyId: 1},
	models.Year{Name: "سال مالی 1397", CompanyId: 1},
	models.Year{Name: "سال مالی 1398", CompanyId: 1},
	models.Year{Name: "سال مالی 1399", CompanyId: 1},
}

var companytypes = []models.CompanyType{
	models.CompanyType{Name: "بازرگانی"},
	models.CompanyType{Name: "تولیدی"},
	models.CompanyType{Name: "خدماتی"},
	models.CompanyType{Name: "پیمانکاری"},
}

var company = []models.Company{
	{
		Name:          "پاک تدبیر موکریان",
		Image:         "",
		CompanyTypeID: 1,
	},
}
