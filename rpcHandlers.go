package main

import (
	"context"
	cutter_proto "links_cutter/api"
	"links_cutter/commands"
)

func (s *server) Create(ctx context.Context, in *cutter_proto.UrlRequest) (*cutter_proto.UrlResponse, error) {
	shortURL := commands.CreateURL(in.Url)

	return &cutter_proto.UrlResponse{
		Url: shortURL,
	}, nil
}

func (s *server) Get(ctx context.Context, in *cutter_proto.UrlRequest) (*cutter_proto.UrlResponse, error) {
	longURL := commands.GetURL(in.Url)

	return &cutter_proto.UrlResponse{
		Url: longURL,
	}, nil
}