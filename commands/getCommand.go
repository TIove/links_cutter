package commands

import "links_cutter/database"

func GetURL(shortURL string) (string, error) {
	return database.GetValue(shortURL)
}
