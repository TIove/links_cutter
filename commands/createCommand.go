package commands

import (
	"links_cutter/database"
	"links_cutter/helpers"
)

func CreateURL(longURL string) string {
	 shortURL := helpers.GetUUID()

	 database.SetValue(shortURL, longURL)

	 return shortURL
}
