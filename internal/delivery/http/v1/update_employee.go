package v1

import (
	"errors"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/dto"
	service "github.com/Nikita-Mihailuk/smartwayTestTask/internal/service/employee"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func (h *HandlerV1) updateEmployee(ctx fiber.Ctx) error {
	employeeID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid employee id")
	}

	var employee dto.UpdateEmployee
	if err = ctx.Bind().JSON(&employee); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request")
	}

	employee.ID = employeeID

	err = h.employeeService.RefreshEmployee(ctx.Context(), employee)
	if err != nil {
		if errors.Is(err, service.ErrEmployeeExist) {
			return fiber.NewError(fiber.StatusNotFound, "employee already exists")
		}
		if errors.Is(err, service.ErrInvalidCompany) {
			return fiber.NewError(fiber.StatusNotFound, "company with this id does not exist")
		}
		if errors.Is(err, service.ErrInvalidDepartment) {
			return fiber.NewError(fiber.StatusNotFound, "department in this company does not exist")
		}
		if errors.Is(err, service.ErrPassportExist) {
			return fiber.NewError(fiber.StatusNotFound, "this passport already exist")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "internal error")
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
