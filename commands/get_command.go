package commands

import (
	"github.com/TIove/links_cutter/app"
)

func GetURL(shortURL string) (string, error) {
	return app.GetValue(shortURL)
}
