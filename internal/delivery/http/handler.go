package http

import (
	"github.com/gofiber/fiber/v3"
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

	api := router.Group("/api")

	for _, version := range h.versions {
		version.InitRoutes(api)
	}
}
