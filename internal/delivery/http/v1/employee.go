package v1

import "github.com/gofiber/fiber/v3"

func (h *HandlerV1) RegisterEmployeeRouts(v1 fiber.Router) {
	userGroup := v1.Group("/employees")

	userGroup.Post("", h.createEmployee)
	userGroup.Get("/company/:id/department/:id", h.getEmployeeByDepartment)
	userGroup.Get("/company/:id", h.getEmployeesByCompany)
	userGroup.Put("/:id", h.updateEmployee)
	userGroup.Delete("/:id", h.deleteEmployee)
}

func (h *HandlerV1) createEmployee(ctx fiber.Ctx) error {
	panic("implement me")
}

func (h *HandlerV1) getEmployeesByCompany(ctx fiber.Ctx) error {
	panic("implement me")
}

func (h *HandlerV1) getEmployeeByDepartment(ctx fiber.Ctx) error {
	panic("implement me")
}

func (h *HandlerV1) updateEmployee(ctx fiber.Ctx) error {
	panic("implement me")
}

func (h *HandlerV1) deleteEmployee(ctx fiber.Ctx) error {
	panic("implement me")
}
