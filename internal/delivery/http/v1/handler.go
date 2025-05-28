package v1

import (
	"context"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/dto"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/model"
	"github.com/gofiber/fiber/v3"
)

type HandlerV1 struct {
	employeeService EmployeeService
}

func NewHandlerV1(employeeService EmployeeService) *HandlerV1 {
	return &HandlerV1{
		employeeService: employeeService,
	}
}

type EmployeeService interface {
	CreateEmployee(ctx context.Context, employee model.Employee) (int, error)
	GetEmployeesByCompany(ctx context.Context, companyID int) ([]model.Employee, error)
	GetEmployeeByDepartment(ctx context.Context, departmentID, companyID int) ([]model.Employee, error)
	DropEmployee(ctx context.Context, id int) error
	RefreshEmployee(ctx context.Context, employee dto.UpdateEmployee) error
}

func (h *HandlerV1) InitRoutes(api fiber.Router) {
	v1 := api.Group("/v1")
	h.RegisterEmployeeRouts(v1)
}

func (h *HandlerV1) RegisterEmployeeRouts(v1 fiber.Router) {
	userGroup := v1.Group("/employees")

	userGroup.Post("", h.createEmployee)
	userGroup.Get("/company/:companyID/department/:departmentID", h.getEmployeeByDepartment)
	userGroup.Get("/company/:companyID", h.getEmployeesByCompany)
	userGroup.Patch("/:id", h.updateEmployee)
	userGroup.Delete("/:id", h.deleteEmployee)
}
