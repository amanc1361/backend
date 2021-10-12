package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var groupRoutes = []Route{
	Route{
		Uri:     "/group/list",
		Method:  http.MethodPost,
		Handler: controllers.GetGroups,
		IsAuth:  false,
	},
	Route{
		Uri:     "/group/get/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetGroup,
		IsAuth:  true,
	},
	Route{
		Uri:     "/groupsbycode/{code}",
		Method:  http.MethodGet,
		Handler: controllers.GetGroupByCode,
		IsAuth:  false,
	},
	Route{
		Uri:     "/groups",
		Method:  http.MethodPost,
		Handler: controllers.CreateGroup,
		IsAuth:  true,
	},
	Route{
		Uri:     "/groups",
		Method:  http.MethodPut,
		Handler: controllers.UpdateGroup,
		IsAuth:  true,
	},
	Route{
		Uri:     "/groups/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteGroup,
		IsAuth:  true,
	},
	Route{
		Uri:     "/group/lastcode",
		Method:  http.MethodPost,
		Handler: controllers.GetLastGroupcode,
		IsAuth:  true,
	},
}
