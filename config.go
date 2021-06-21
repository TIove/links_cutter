package main

import (
	"log"
	"os"
)

const (
	port     = ":8001"
	grpcPort = ":50051"
)


var app Application

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC)
var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.LUTC|log.Lshortfile)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

