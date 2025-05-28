package v1

import (
	"errors"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/dto"
	service "github.com/Nikita-Mihailuk/smartwayTestTask/internal/service/employee"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// @Summary Update employee information
// @Description Update existing employee data by ID
// @Tags V1
// @Accept json
// @Param id path int true "Employee ID"
// @Param employee body dto.UpdateEmployee true "Updated employee data"
// @Success 204 "No content"
// @Failure 400 {string} string "Invalid request data"
// @Failure 404 {string} string "Employee, company or department not found"
// @Failure 409 {string} string "Passport already exists"
// @Failure 500 {string} string "Internal server error"
// @Router /api/v1/employees/{id} [patch]
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
			return fiber.NewError(fiber.StatusConflict, "employee already exists")
		}
		if errors.Is(err, service.ErrInvalidCompany) {
			return fiber.NewError(fiber.StatusNotFound, "company with this id does not exist")
		}
		if errors.Is(err, service.ErrInvalidDepartment) {
			return fiber.NewError(fiber.StatusNotFound, "department in this company does not exist")
		}
		if errors.Is(err, service.ErrPassportExist) {
			return fiber.NewError(fiber.StatusConflict, "this passport already exist")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "internal error")
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
