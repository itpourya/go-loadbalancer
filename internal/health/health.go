package health

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"github.com/itpourya/go-loadbalancer/internal/server"
	"log/slog"
	"net/http"
	"time"
)

func StartHealthCheck(server []*server.Server) {
	cr, _ := gocron.NewScheduler()

	_, err := cr.NewJob(gocron.DurationJob(2*time.Second), gocron.NewTask(func() {
		for _, host := range server {
			healthy := checkHealth(host)

			if healthy {
				slog.Info(host.Name + " is healthy")
			} else {
				slog.Info(host.Name + " is not healthy")
			}
		}
	}))
	if err != nil {
		fmt.Println(err)
		slog.Error(err.Error())
	}

	cr.Start()

	select {}

	cr.Shutdown()
}

func checkHealth(server *server.Server) bool {
	response, err := http.Head(server.URL)
	if err != nil {
		server.Health = false
		return server.Health
	}

	if response.StatusCode != 200 {
		server.Health = false
		return server.Health
	}

	server.Health = true
	return server.Health
}
