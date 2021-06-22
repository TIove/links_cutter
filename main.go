package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	cutter "links_cutter/api"
	. "links_cutter/app"
	"net"
	"os"
)

type server struct {
	cutter.UnimplementedCutterServer
}

func main() {
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbCredentials := fmt.Sprintf("user=%s password=%s dbname=postgres sslmode=disable", user, password)
	db, err := openDb("postgres", dbCredentials)

	if err != nil {
		ErrorLog.Fatal(err)
	}
	defer db.Close()

	App = Application{
		ErrorLog: ErrorLog,
		InfoLog:  InfoLog,
		Db:       &DbModel{Db: db},
	}

	upGrpcServer()
}

func upGrpcServer() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		App.ErrorLog.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	cutter.RegisterCutterServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		App.ErrorLog.Fatalf("Failed to serve: %v", err)
	}
}

func openDb(driverName string, dbCredentials string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dbCredentials)
	if err != nil {
		panic(err)
	}

	return db, err
}
