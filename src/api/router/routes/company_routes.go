package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var companyRoutes = []Route{
	Route{
		Uri:     "/company",
		Method:  http.MethodGet,
		Handler: controllers.GetCompanys,
		IsAuth:  false,
	},
	Route{
		Uri:     "/company/{id}",
		Method:  http.MethodGet,
		
		Handler: controllers.GetCompany,
		IsAuth:  false,
	},
	Route{
		Uri:     "/company",
		Method:  http.MethodPost,
		Handler: controllers.CreateCompany,
		IsAuth:  true,
	},
	Route{
		Uri:     "/company/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateCompany,
		IsAuth:  true,
	},
	Route{
		Uri:     "/company/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteCompany,
		IsAuth:  true,
	},
}
