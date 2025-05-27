package main

import (
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/app"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/config"
	"github.com/Nikita-Mihailuk/smartwayTestTask/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	log := logging.GetLogger(cfg.Env)

	application := app.NewApp(log, cfg)
	application.HTTPServer.MustRun()
}
