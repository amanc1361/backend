package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var costcenterRoutes = []Route{
	Route{
		Uri:     "/costcenter",
		Method:  http.MethodGet,
		Handler: controllers.GetCostCenters,
		IsAuth:  true,
	},
	Route{
		Uri:     "/costcenterbyid",
		Method:  http.MethodPost,
		
		Handler: controllers.GetCostCenterByID,
		IsAuth:  true,
	},
	Route{
		Uri:     "/costcenter",
		Method:  http.MethodPost,
		Handler: controllers.CreateCostCenter,
		IsAuth:  true,
	},
	Route{
		Uri:     "/costcenter",
		Method:  http.MethodPut,
		Handler: controllers.UpdateCostCenter,
		IsAuth:  true,
	},
	Route{
		Uri:     "/costcenter",
		Method:  http.MethodDelete,
		Handler: controllers.CostDelete,
		IsAuth:  true,
	},
}
