package v1

import (
	"errors"
	service "github.com/Nikita-Mihailuk/smartwayTestTask/internal/service/employee"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func (h *HandlerV1) getEmployeesByCompany(ctx fiber.Ctx) error {
	companyID, err := strconv.Atoi(ctx.Params("companyID"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid company id")
	}

	employees, err := h.employeeService.GetEmployeesByCompany(ctx.Context(), companyID)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCompany) {
			return fiber.NewError(fiber.StatusNotFound, "company with this id does not exist")
		}
		if errors.Is(err, service.ErrEmployeesNotFoundByCompany) {
			return fiber.NewError(fiber.StatusNotFound, "employees not found by company")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "internal error")
	}

	return ctx.JSON(employees)
}
