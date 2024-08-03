package main

import (
	"github.com/itpourya/go-loadbalancer/internal/handler"
	"github.com/itpourya/go-loadbalancer/internal/health"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler.ForwardRequest)
	slog.NewJSONHandler(os.Stdout, nil)

	go health.StartHealthCheck(handler.ServerList)

	slog.Info("Start Load Balancer PORT:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
