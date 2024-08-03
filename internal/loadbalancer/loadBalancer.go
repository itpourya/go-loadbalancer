package loadbalancer

import (
	"github.com/itpourya/go-loadbalancer/internal/server"
	"log/slog"
)

var (
	lastServedIndex = 0
)

func getServer(serverListCount int) int {
	nextIndex := (lastServedIndex + 1) % serverListCount
	lastServedIndex = nextIndex

	return nextIndex
}

func GetHealthyServer(serverList []*server.Server) *server.Server {
	for i := 0; i < len(serverList); i++ {
		mark := getServer(len(serverList))
		if serverList[mark].Health {
			return serverList[mark]
		}
	}

	slog.Warn("No healthy server exist")
	return nil
}
