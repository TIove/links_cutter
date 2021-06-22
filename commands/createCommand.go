package commands

import (
	"links_cutter/app"
	"links_cutter/helpers"
)

func CreateURL(longURL string) string {
	shortURL := helpers.GetUUID()

	 app.SetValue(shortURL, longURL)

	return shortURL
}
