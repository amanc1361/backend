package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var yearRoutes = []Route{
	Route{
		Uri:     "/years",
		Method:  http.MethodGet,
		Handler: controllers.GetYears,
		IsAuth:  false,
	},
	Route{
		Uri:     "/years/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetYear,
		IsAuth:  false,
	},
	Route{
		Uri:     "/years",
		Method:  http.MethodPost,
		Handler: controllers.CreateYear,
		IsAuth:  true,
	},
	Route{
		Uri:     "/years",
		Method:  http.MethodPut,
		Handler: controllers.UpdateYear,
		IsAuth:  true,
	},
	Route{
		Uri:     "/years/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteYear,
		IsAuth:  true,
	},
}
