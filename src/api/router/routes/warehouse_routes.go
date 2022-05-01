package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var wareHouseRoute = []Route{
	{
		Uri:     "/warehouse",
		Method:  http.MethodGet,
		Handler: controllers.GetWareHousings,
		IsAuth:  true,
	},

}