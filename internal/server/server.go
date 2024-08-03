package server

import (
	"log/slog"
	"net/http/httputil"
	url2 "net/url"
)

type Server struct {
	Name         string
	URL          string
	ReverseProxy *httputil.ReverseProxy
	Health       bool
}

func NewServer(name, url string) *Server {
	urlParse, err := url2.Parse(url)
	if err != nil {
		slog.Error("Can not parse url")
		return nil
	}

	return &Server{
		Name:         name,
		URL:          url,
		ReverseProxy: httputil.NewSingleHostReverseProxy(urlParse),
	}
}
