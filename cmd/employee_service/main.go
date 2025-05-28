package main

import (
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/app"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/config"
	"github.com/Nikita-Mihailuk/smartwayTestTask/pkg/logging"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.GetConfig()
	log := logging.GetLogger(cfg.Env)

	application := app.NewApp(log, cfg)
	go application.HTTPServer.MustRun()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	application.HTTPServer.Stop()
	log.Info("application stopped")
}
