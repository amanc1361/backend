package middlewares

import (
	"back-account/src/api/auth"
	"back-account/src/api/models"
	"back-account/src/api/responses"
	"back-account/src/api/utils/types"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Printf("%s %s %s%s", r.Host, r.RequestURI, r.Method, r.Proto)
		next(w, r)
	}
}

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		next(w, r)
	}
}

// CORS Middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("ok")

		// Next
		next.ServeHTTP(w, r)
		return
	})
}
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleware", r.Method)

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Content-Type", "application/json")
			return
		}

		next.ServeHTTP(w, r)
		log.Println("Executing middleware again")
	})
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := auth.ExtractToken(w, r)
		if token == nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("توکن شما منقضی است"))
			return
		}
		if token.Valid {
			ctx := context.WithValue(
				r.Context(),
				types.UserKey("user"),
				token.Claims.(*models.Claim).User,
			)
			next(w, r.WithContext(ctx))
		}
	}
}
