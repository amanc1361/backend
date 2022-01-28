package routes

import (
	"back-account/src/api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
	IsAuth  bool
}

func Load() []Route {
	routes := userRoutes
	routes = append(routes, yearRoutes...)
	routes = append(routes, companyTypeRoutes...)
	routes = append(routes, companyRoutes...)
	routes = append(routes, groupRoutes...)
	routes = append(routes, ledgerRoutes...)
	routes = append(routes, subledgerRoutes...)
	routes = append(routes, detailedRoutes...)
	routes = append(routes, documentRoutes...)
	routes = append(routes, loginRoutes...)
	routes = append(routes, bilanRoutes...)
	routes = append(routes, storeRoutes...)
	routes = append(routes, storeGroupRoutes...)
	routes = append(routes, storeSubGroupRoutes...)
	routes = append(routes, storechildsubgroupRoutes...)
	routes = append(routes, storeObjectRoutes...)
	routes = append(routes, storeActionRoutes...)
	routes = append(routes, unitRoutes...)
	routes = append(routes, storeactiontypeRoutes...)
	routes = append(routes, storepersonRoutes...)
	routes = append(routes, storepersonreciveRoutes...)
	routes = append(routes, peopleRoutes...)
	routes = append(routes, documentmergeroutes...)
	routes=append(routes, costcenterRoutes...)

	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {

	for _, route := range Load() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)

	}
	return r
}

func SetupRoutesWithMiddlware(r *mux.Router) *mux.Router {

	for _, route := range Load() {
		if route.IsAuth {
			r.HandleFunc(

				route.Uri,

				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(
						middlewares.SetMiddlewareAuthentication(route.Handler),
					),
				),
			).Methods(route.Method, http.MethodOptions)
		} else {
			r.HandleFunc(route.Uri,

				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(route.Handler),
				),
			).Methods(route.Method, http.MethodOptions)
		}

	}

	return r

}
