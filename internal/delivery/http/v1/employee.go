package v1

import (
	"errors"
	"fmt"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/dto"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/model"
	service "github.com/Nikita-Mihailuk/smartwayTestTask/internal/service/employee"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func (h *HandlerV1) RegisterEmployeeRouts(v1 fiber.Router) {
	userGroup := v1.Group("/employees")

	userGroup.Post("", h.createEmployee)
	userGroup.Get("/company/:companyID/department/:departmentID", h.getEmployeeByDepartment)
	userGroup.Get("/company/:companyID", h.getEmployeesByCompany)
	userGroup.Patch("/:id", h.updateEmployee)
	userGroup.Delete("/:id", h.deleteEmployee)
}

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
		if errors.Is(err, service.ErrDepartmentExist) {
			return fiber.NewError(fiber.StatusConflict, "department with number already exists")
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

func validateCreateEmployee(employee model.Employee) error {
	if employee.Name == "" || employee.Surname == "" ||
		employee.Phone == "" || employee.CompanyID == 0 ||
		employee.Department.Phone == "" || employee.Department.Name == "" ||
		employee.Passport.Type == "" || employee.Passport.Number == "" {

		return fmt.Errorf("invalid employee parameters")
	}
	return nil
}
