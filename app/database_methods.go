package app

import (
	"database/sql"
	"github.com/TIove/links_cutter/models"
)

type DbModel struct {
	Db *sql.DB
}

func (model *DbModel) Insert(linkRequest models.Link) (*models.Link, error) {
	queryCheck := "SELECT * FROM links WHERE longurl = $1 LIMIT 1"

	existingData := model.Db.QueryRow(queryCheck, linkRequest.LongURL)
	result := models.Link{}
	err := existingData.Scan(&result.ShortURL, &result.LongURL, &result.CreatedAt)

	if err == nil {
		return &result, err
	}

	queryCheck = "SELECT * FROM links WHERE shorturl = $1 LIMIT 1"

	existingData = model.Db.QueryRow(queryCheck, linkRequest.LongURL)
	result = models.Link{}
	err = existingData.Scan(&result.ShortURL, &result.LongURL, &result.CreatedAt)

	if err == nil {
		ErrorLog.Print("Try again, server error")
		return nil, err
	}

	query := "INSERT INTO links VALUES ($1, $2, $3)"

	_, err = model.Db.Exec(query, linkRequest.ShortURL, linkRequest.LongURL, linkRequest.CreatedAt)
	return nil, err
}

func (model *DbModel) Get(shortURL string) (*models.Link, error) {
	query := "select * from links where shorturl = $1 LIMIT 1"

	row := model.Db.QueryRow(query, shortURL)

	result := models.Link{}
	err := row.Scan(&result.ShortURL, &result.LongURL, &result.CreatedAt)

	return &result, err
}
