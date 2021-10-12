package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var companyTypeRoutes = []Route{
	Route{
		Uri:     "/companytypes",
		Method:  http.MethodGet,
		Handler: controllers.GetCompanyTypes,
		IsAuth:  false,
	},
	Route{
		Uri:     "/companytypes/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetCompanyType,
		IsAuth:  false,
	},
}
