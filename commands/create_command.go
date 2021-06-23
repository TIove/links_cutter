package commands

import (
	"github.com/TIove/links_cutter/app"
	"github.com/TIove/links_cutter/helpers"
)

func CreateURL(longURL string) string {
	shortURL := helpers.GetUUID()

	app.SetValue(shortURL, longURL)

	return shortURL
}
