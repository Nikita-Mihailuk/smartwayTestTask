package v1

import (
	"errors"
	service "github.com/Nikita-Mihailuk/smartwayTestTask/internal/service/employee"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func (h *HandlerV1) getEmployeeByDepartment(ctx fiber.Ctx) error {
	companyID, err := strconv.Atoi(ctx.Params("companyID"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid company id")
	}

	departmentID, err := strconv.Atoi(ctx.Params("departmentID"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid department id")
	}

	employees, err := h.employeeService.GetEmployeeByDepartment(ctx.Context(), companyID, departmentID)
	if err != nil {
		if errors.Is(err, service.ErrInvalidDepartment) {
			return fiber.NewError(fiber.StatusNotFound, "department in this company does not exist")
		}
		if errors.Is(err, service.ErrEmployeesNotFoundByDepartment) {
			return fiber.NewError(fiber.StatusNotFound, "employees not found by department")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "internal error")
	}

	return ctx.JSON(employees)
}
