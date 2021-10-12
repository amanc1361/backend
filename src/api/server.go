package api

import (
	"back-account/src/api/auto"
	"back-account/src/api/config"

	"back-account/src/api/router"

	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func Run() {

	config.Load()
	auto.Load()
	fmt.Printf("\n\tListining[::]%d\n", config.PORT)

	listen(config.PORT)
}
func listen(port int) {
	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},                            // All origins
	// 	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Allowing only get, just an example
	// })
	r := router.New()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), cors.AllowAll().Handler(r)))
}
