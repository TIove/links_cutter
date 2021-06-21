package main

import (
	"links_cutter/commands"
	"net/http"
)

func (app *Application) get(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	shortUrl := r.URL.Query().Get("url")

	longUrl := commands.GetURL(shortUrl)

	_, err := w.Write([]byte(longUrl))
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
}

func (app *Application) create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	longURL := r.URL.Query().Get("url")

	shortURL := commands.CreateURL(longURL)

	_, err := w.Write([]byte(shortURL))
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
}
