package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFoundResponse(w, r)
	})

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/list/:id", app.listHandlerGet)
	router.HandlerFunc(http.MethodPost, "/v1/list", app.listHandlerPost)
	router.HandlerFunc(http.MethodGet, "/v1/task/:id", app.taskHandlerGet)
	router.HandlerFunc(http.MethodPost, "/v1/task", app.taskHandlerPost)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}
