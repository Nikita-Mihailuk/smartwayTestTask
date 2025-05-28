package http

import (
	_ "github.com/Nikita-Mihailuk/smartwayTestTask/docs"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	versions []APIDelivery
}

type APIDelivery interface {
	InitRoutes(router fiber.Router)
}

func NewHandler(versions ...APIDelivery) *Handler {
	return &Handler{versions: versions}
}

func (h *Handler) InitHandler(router fiber.Router) {
	router.Get("/ping", func(ctx fiber.Ctx) error {
		return ctx.SendString("pong")
	})
	router.Get("/swagger/*", SwaggerHandler())

	api := router.Group("/api")

	for _, version := range h.versions {
		version.InitRoutes(api)
	}
}

func SwaggerHandler() fiber.Handler {
	return adaptor.HTTPHandler(httpSwagger.WrapHandler)
}
