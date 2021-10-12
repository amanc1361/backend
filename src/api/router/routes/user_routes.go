package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var userRoutes = []Route{
	Route{
		Uri:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
		IsAuth:  false,
	},
	Route{
		Uri:     "/users/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetUser,
		IsAuth:  false,
	},
	Route{
		Uri:     "/users",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
		IsAuth:  true,
	},
	Route{
		Uri:     "/users/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateUser,
		IsAuth:  true,
	},
	Route{
		Uri:     "/users/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteUser,
		IsAuth:  true,
	},
}
