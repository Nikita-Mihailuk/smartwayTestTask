package v1

import (
	"errors"
	"fmt"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/model"
	service "github.com/Nikita-Mihailuk/smartwayTestTask/internal/service/employee"
	"github.com/gofiber/fiber/v3"
)

func (h *HandlerV1) createEmployee(ctx fiber.Ctx) error {
	var employee model.Employee

	if err := ctx.Bind().JSON(&employee); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request")
	}

	if err := validateCreateEmployee(employee); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id, err := h.employeeService.CreateEmployee(ctx.Context(), employee)
	if err != nil {
		if errors.Is(err, service.ErrEmployeeExist) {
			return fiber.NewError(fiber.StatusConflict, "employee with number already exists")
		}
		if errors.Is(err, service.ErrInvalidDepartment) {
			return fiber.NewError(fiber.StatusConflict, "department with this id does not exist")
		}
		if errors.Is(err, service.ErrInvalidCompany) {
			return fiber.NewError(fiber.StatusNotFound, "company with this id does not exist")
		}
		if errors.Is(err, service.ErrPassportExist) {
			return fiber.NewError(fiber.StatusConflict, "this passport already exist")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "internal error")
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

func validateCreateEmployee(employee model.Employee) error {
	if employee.Name == "" || employee.Surname == "" ||
		employee.Phone == "" || employee.CompanyID == 0 ||
		employee.Department.Phone == "" || employee.Department.Name == "" ||
		employee.Passport.Type == "" || employee.Passport.Number == "" {

		return fmt.Errorf("invalid employee parameters")
	}
	return nil
}
