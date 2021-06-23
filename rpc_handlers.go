package main

import (
	"context"
	cutterProto "github.com/TIove/links_cutter/api"
	"github.com/TIove/links_cutter/commands"
)

func (s *server) Create(_ context.Context, in *cutterProto.UrlRequest) (*cutterProto.UrlResponse, error) {
	shortURL := commands.CreateURL(in.Url)

	return &cutterProto.UrlResponse{
		Url: shortURL,
	}, nil
}

func (s *server) Get(_ context.Context, in *cutterProto.UrlRequest) (*cutterProto.UrlResponse, error) {
	longURL, err := commands.GetURL(in.Url)

	return &cutterProto.UrlResponse{
		Url: longURL,
	}, err
}
