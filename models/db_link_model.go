package models

import "time"

type Link struct {
	ShortURL  string
	LongURL   string
	CreatedAt time.Time
}
