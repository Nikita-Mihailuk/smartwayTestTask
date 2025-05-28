package v1

import (
	"errors"
	service "github.com/Nikita-Mihailuk/smartwayTestTask/internal/service/employee"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// @Summary Get employees by company and department
// @Description Retrieve employees belonging to specific company and department
// @Tags V1
// @Produce json
// @Param companyID path int true "Company ID"
// @Param departmentID path int true "Department ID"
// @Success 200 {array} model.Employee "List of employees"
// @SuccessExample {json} Success-Response:
//
//	[{
//	    "name": "test",
//	    "surname": "test",
//	    "phone": "+7-202-555-0111",
//	    "company_id": 1,
//	    "passport": {
//	        "type": "3333",
//	        "number": "111121"
//	    },
//	    "department": {
//	        "name": "IT",
//	        "phone": "+7 900 123 33 11"
//	    }
//	}]
//
// @Failure 400 {string} string "Invalid company or department ID"
// @Failure 404 {string} string "Department or employees not found"
// @Failure 500 {string} string "Internal server error"
// @Router /api/v1/employees/company/{companyID}/department/{departmentID} [get]
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
