package router

import (
	"back-account/src/api/router/routes"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	//  r.Use(middlewares.CORS)
	// r.Use(mux.CORSMethodMiddleware(r))
	return routes.SetupRoutesWithMiddlware(r)
}
