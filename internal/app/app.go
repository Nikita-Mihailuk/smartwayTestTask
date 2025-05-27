package app

import (
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/app/http"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/config"
	v1 "github.com/Nikita-Mihailuk/smartwayTestTask/internal/delivery/http/v1"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/service/employee"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/storage/postgres"
	"go.uber.org/zap"
)

type App struct {
	HTTPServer *http.App
}

func NewApp(log *zap.Logger, cfg *config.Config) *App {

	storage := postgres.NewStorage(cfg)

	service := employee.NewEmployeeService(log, storage, storage, storage, storage)

	handlerV1 := v1.NewHandlerV1(service)

	httpApp := http.NewApp(cfg.Server.Port, cfg.Server.Host, handlerV1)

	return &App{
		HTTPServer: httpApp,
	}
}
