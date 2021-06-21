package main

import (
	"links_cutter/commands"
	"net/http"
)

func (app *Application) get(w http.ResponseWriter, r *http.Request) { // Add get
	shortUrl := r.URL.Query().Get("url")

	longUrl := commands.GetURL(shortUrl)

	w.Write([]byte(longUrl))
}

func (app *Application) create(w http.ResponseWriter, r *http.Request) { // Add post
	longURL := r.URL.Query().Get("url")

	shortURL := commands.CreateURL(longURL)

	w.Write([]byte(shortURL))
}
