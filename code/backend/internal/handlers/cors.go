package handlers

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func NewCors() func(http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedMethods([]string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
		}),
		handlers.AllowedOrigins([]string{
			"http://localhost:3000",
		}),
		handlers.AllowedHeaders([]string{
			"auth-x",
			"Content-Type",
		}),
		handlers.ExposedHeaders([]string{
			"x-auth",
		}),
		handlers.AllowCredentials(),
	)
}
