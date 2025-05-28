package v1

import (
	"errors"
	service "github.com/Nikita-Mihailuk/smartwayTestTask/internal/service/employee"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// @Summary Delete employee
// @Description Delete employee by ID
// @Tags V1
// @Param id path int true "Employee ID"
// @Success 204 "No content"
// @Failure 400 {string} string "Invalid employee ID"
// @Failure 404 {string} string "Employee not found"
// @Failure 500 {string} string "Internal server error"
// @Router /api/v1/employees/{id} [delete]
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
