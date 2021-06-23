package app

import (
	"log"
	"os"
)

var App Application

var InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC)
var ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.LUTC|log.Lshortfile)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Db       *DbModel
}

