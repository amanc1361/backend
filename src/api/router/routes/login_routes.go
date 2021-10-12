package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var loginRoutes = []Route{
	Route{
		Uri:     "/login",
		Method:  http.MethodPost,
		Handler: controllers.Login,
		IsAuth:  false,
	},
}
