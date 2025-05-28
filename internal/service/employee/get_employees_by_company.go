package employee

import (
	"context"
	"errors"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/model"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/storage/postgres"
	"go.uber.org/zap"
)

func (s *EmployeeService) GetEmployeesByCompany(ctx context.Context, companyID int) ([]model.Employee, error) {
	employees, err := s.employeeProvider.GetEmployeesByCompanyID(ctx, companyID)
	if err != nil {
		if errors.Is(err, postgres.ErrCompanyNotFound) {
			s.log.Error("company not found", zap.Error(err))
			return nil, ErrInvalidCompany
		}
		if errors.Is(err, postgres.ErrEmployeesNotFound) {
			s.log.Error("employees not found by company", zap.Error(err))
			return nil, ErrEmployeesNotFoundByCompany
		}
		s.log.Error("failed to fetch employees by company", zap.Error(err))
		return nil, err
	}

	s.log.Info("fetched employees by company", zap.Int("count", len(employees)))
	return employees, nil
}
