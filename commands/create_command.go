package commands

import (
	"github.com/TIove/links_cutter/app"
	"github.com/TIove/links_cutter/helpers"
)

func CreateURL(longURL string) string {
	shortURL := helpers.GetUUID()

	oldUrl := app.SetValue(shortURL, longURL)
	if oldUrl != "" {
		shortURL = oldUrl
	}

	return shortURL
}
