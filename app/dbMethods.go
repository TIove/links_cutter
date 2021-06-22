package app

import (
	"database/sql"
	"links_cutter/models"
)

type DbModel struct {
	Db *sql.DB
}

func (model *DbModel) Insert(linkRequest models.Link) (sql.Result, error) {
	query := "INSERT INTO links VALUES ($1, $2, $3)"

	res, err := model.Db.Exec(query, linkRequest.ShortURL, linkRequest.LongURL, linkRequest.CreatedAt)
	return res, err
}

func (model *DbModel) Get(shortURL string) (*models.Link, error) { // TODO is pointer or not?
	query := "select * from users where id = $1 LIMIT 1"

	row := model.Db.QueryRow(query, shortURL)

	result := models.Link{}
	err := row.Scan(&result.ShortURL, &result.LongURL, &result.CreatedAt)

	return &result, err
}