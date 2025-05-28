package v1

import (
	"errors"
	service "github.com/Nikita-Mihailuk/smartwayTestTask/internal/service/employee"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// @Summary Get employees by company ID
// @Description Retrieve all employees belonging to a specific company
// @Tags V1
// @Produce json
// @Param companyID path int true "Company ID"
// @Success 200 {array} model.Employee "List of employees"
// @SuccessExample {json} Success-Response:
//
//	[{
//	    "name": "Noah",
//	    "surname": "Clark",
//	    "phone": "+7-202-555-0114",
//	    "company_id": 1,
//	    "passport": {
//	        "type": "3333",
//	        "number": "111111"
//	    },
//	    "department": {
//	        "name": "Support",
//	        "phone": "89001003243"
//	    }
//	}]
//
// @Failure 400 {string} string "Invalid company ID"
// @Failure 404 {string} string "Company or employees not found"
// @Failure 500 {string} string "Internal server error"
// @Router /api/v1/employees/company/{companyID} [get]
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
