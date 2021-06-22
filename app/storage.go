package app

import (
	"links_cutter/models"
	"log"
	"time"
)

type AppDb struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Db       *DbModel
}

//var urls = make(map[string]string)

func GetValue(key string) (string, error) {
	val, err := App.Db.Get(key)
	if err != nil {
		App.ErrorLog.Println(err)
		return "", err
	}

	result := val.LongURL

	return result, nil
}

func SetValue(key string, value string) {
	request := models.Link{ShortURL: key, LongURL: value, CreatedAt: time.Now()}

	_, err := App.Db.Insert(request)
	if err != nil {
		App.ErrorLog.Print(err)
		return
	}
}
