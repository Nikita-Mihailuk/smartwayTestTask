package employee

import (
	"context"
	"errors"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/model"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/storage/postgres"
	"go.uber.org/zap"
)

func (s *EmployeeService) GetEmployeeByDepartment(ctx context.Context, departmentID, companyID int) ([]model.Employee, error) {
	employees, err := s.employeeProvider.GetEmployeeByDepartmentID(ctx, departmentID, companyID)
	if err != nil {
		if errors.Is(err, postgres.ErrDepartmentNotFound) {
			s.log.Error("department not found", zap.Error(err))
			return nil, ErrInvalidDepartment
		}
		if errors.Is(err, postgres.ErrEmployeesNotFound) {
			s.log.Error("employees not found by department", zap.Error(err))
			return nil, ErrEmployeesNotFoundByDepartment
		}
		s.log.Error("failed to fetch employees by department", zap.Error(err))
		return nil, err
	}

	s.log.Info("fetched employees by department", zap.Int("count", len(employees)))
	return employees, nil
}
