package app

import (
	"github.com/TIove/links_cutter/models"
	"log"
	"time"
)

type Db struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Db       *DbModel
}

func GetValue(key string) (string, error) {
	val, err := App.Db.Get(key)
	if err != nil {
		App.ErrorLog.Println(err)
		return "", err
	}

	result := val.LongURL

	return result, nil
}

func SetValue(key string, value string) string {
	request := models.Link{
		ShortURL:  key,
		LongURL:   value,
		CreatedAt: time.Now(),
	}

	res, err := App.Db.Insert(request)
	if res != nil {
		return res.ShortURL
	}
	if err != nil {
		App.ErrorLog.Print(err)
	}
	return ""
}
