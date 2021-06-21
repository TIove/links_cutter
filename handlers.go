package main

import (
	"links_cutter/commands"
	"links_cutter/helpers"
	"net/http"
	"strings"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	path := strings.Split(r.URL.Path, "/")
	shortUrl := path[len(path)-1]

	longUrl, err := commands.GetURL(shortUrl)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		app.ErrorLog.Print(err)
		return
	}

	if err := helpers.OpenBrowser(longUrl); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		app.ErrorLog.Print(err)
		return
	}
}
