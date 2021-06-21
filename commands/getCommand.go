package commands

import "links_cutter/database"

func GetURL(shortURL string) string {
	return database.GetValue(shortURL)
}