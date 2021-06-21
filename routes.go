package main

import (
	"net/http"
)

func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	userRoutes(app, mux)

	return mux
}

func userRoutes(app *Application, mux *http.ServeMux) {
	mux.HandleFunc("/", app.home)
}
