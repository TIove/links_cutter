package commands

import (
	"links_cutter/app"
)

func GetURL(shortURL string) (string, error) {
	return app.GetValue(shortURL)
}
