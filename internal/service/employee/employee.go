package employee

import (
	"context"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/model"
)

func (s *EmployeeService) CreateEmployee(ctx context.Context, employee model.Employee) (int, error) {
	panic("implement me")
}

func (s *EmployeeService) GetEmployeesByCompany(ctx context.Context, companyID int) ([]model.Employee, error) {
	panic("implement me")
}
func (s *EmployeeService) GetEmployeeByDepartment(ctx context.Context, departmentID, companyID int) ([]model.Employee, error) {
	panic("implement me")
}
func (s *EmployeeService) DropEmployee(ctx context.Context, id int) error {
	panic("implement me")
}
func (s *EmployeeService) RefreshEmployee(ctx context.Context, employee model.Employee) error {
	panic("implement me")
}
