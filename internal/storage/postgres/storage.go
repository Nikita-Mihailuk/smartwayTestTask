package postgres

import (
	"context"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/model"
)

func (s *Storage) SaveEmployee(ctx context.Context, employee model.Employee) (int, error) {
	panic("implement me")
}

func (s *Storage) GetEmployeesByCompanyID(ctx context.Context, companyID int) ([]model.Employee, error) {
	panic("implement me")
}

func (s *Storage) GetEmployeeByDepartmentID(ctx context.Context, departmentID, companyID int) ([]model.Employee, error) {
	panic("implement me")
}

func (s *Storage) DeleteEmployee(ctx context.Context, id int) error {
	panic("implement me")
}

func (s *Storage) UpdateEmployee(ctx context.Context, employee model.Employee) error {
	panic("implement me")
}
