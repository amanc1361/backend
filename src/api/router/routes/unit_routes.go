package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var unitRoutes = []Route{
	Route{
		Uri:     "/unit/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetUnit,
		IsAuth:  true,
	},
	Route{
		Uri:     "/units",
		Method:  http.MethodPost,
		Handler: controllers.GetUnits,
		IsAuth:  true,
	},

	Route{
		Uri:     "/unit",
		Method:  http.MethodPost,
		Handler: controllers.CreateUnit,
		IsAuth:  true,
	},
	Route{
		Uri:     "/unit/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateUnit,
		IsAuth:  true,
	},
	Route{
		Uri:     "/unit/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteUnit,
		IsAuth:  true,
	},
}
