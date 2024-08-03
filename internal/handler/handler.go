package handler

import (
	"github.com/itpourya/go-loadbalancer/internal/loadbalancer"
	"github.com/itpourya/go-loadbalancer/internal/server"
	"net/http"
)

var (
	ServerList = []*server.Server{
		server.NewServer("server-1", "http://127.0.0.1:8080"),
		server.NewServer("server-2", "http://127.0.0.1:8081"),
		server.NewServer("server-3", "http://127.0.0.1:8082"),
		server.NewServer("server-4", "http://127.0.0.1:8083"),
		server.NewServer("server-5", "http://127.0.0.1:8084"),
	}
)

func ForwardRequest(res http.ResponseWriter, req *http.Request) {
	srv := loadbalancer.GetHealthyServer(ServerList)

	srv.ReverseProxy.ServeHTTP(res, req)
}
