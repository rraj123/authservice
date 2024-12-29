package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/auth/", app.getAuthCode)
	router.HandlerFunc(http.MethodGet, "/v1/reflect/", app.getRequestInfo)

	// Return the httprouter instance.
	return router
}
