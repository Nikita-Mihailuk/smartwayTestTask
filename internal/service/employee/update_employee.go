package employee

import (
	"context"
	"errors"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/dto"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/storage/postgres"
	"go.uber.org/zap"
)

func (s *EmployeeService) RefreshEmployee(ctx context.Context, employee dto.UpdateEmployee) error {
	err := s.employeeUpdater.UpdateEmployee(ctx, employee)
	if err != nil {
		if errors.Is(err, postgres.ErrEmployeeExist) {
			s.log.Error("employee already exists", zap.Error(err))
			return ErrEmployeeExist
		}
		if errors.Is(err, postgres.ErrCompanyNotFound) {
			s.log.Error("company not found", zap.Error(err))
			return ErrInvalidCompany
		}
		if errors.Is(err, postgres.ErrDepartmentNotFound) {
			s.log.Error("department not found", zap.Error(err))
			return ErrInvalidDepartment
		}
		if errors.Is(err, postgres.ErrPassportExist) {
			s.log.Error("passport already exists", zap.Error(err))
			return ErrPassportExist
		}
		s.log.Error("failed to update employee", zap.Error(err))
		return err
	}
	s.log.Info("updated employee", zap.Int("id", employee.ID))
	return nil
}
