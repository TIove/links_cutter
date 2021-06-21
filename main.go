package main

import (
	"google.golang.org/grpc"
	cutter "links_cutter/api"
	"net"
	"net/http"
)

type server struct {
	cutter.UnimplementedCutterServer
}

func main() {
	app = Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	go upGrpcServer()

	err := srv.ListenAndServe()
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
}

func upGrpcServer() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		app.ErrorLog.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	cutter.RegisterCutterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		app.ErrorLog.Fatalf("Failed to serve: %v", err)
	}
}
