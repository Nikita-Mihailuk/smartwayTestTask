package v1

import (
	"errors"
	service "github.com/Nikita-Mihailuk/smartwayTestTask/internal/service/employee"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func (h *HandlerV1) deleteEmployee(ctx fiber.Ctx) error {
	employeeID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid employee id")
	}

	err = h.employeeService.DropEmployee(ctx.Context(), employeeID)
	if err != nil {
		if errors.Is(err, service.ErrEmployeeNotFoundByID) {
			return fiber.NewError(fiber.StatusNotFound, "employee with id does not exist")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "internal error")
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
