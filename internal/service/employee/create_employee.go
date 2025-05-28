package employee

import (
	"context"
	"errors"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/model"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/storage/postgres"
	"go.uber.org/zap"
)

func (s *EmployeeService) CreateEmployee(ctx context.Context, employee model.Employee) (int, error) {
	id, err := s.employeeSaver.SaveEmployee(ctx, employee)
	if err != nil {
		if errors.Is(err, postgres.ErrCompanyNotFound) {
			s.log.Error("company not found", zap.Error(err))
			return 0, ErrInvalidCompany
		}
		if errors.Is(err, postgres.ErrEmployeeExist) {
			s.log.Error("employee already exists", zap.Error(err))
			return 0, ErrEmployeeExist
		}
		if errors.Is(err, postgres.ErrDepartmentNotFound) {
			s.log.Error("department not found", zap.Error(err))
			return 0, ErrInvalidDepartment
		}
		if errors.Is(err, postgres.ErrPassportExist) {
			s.log.Error("passport already exists", zap.Error(err))
			return 0, ErrPassportExist
		}
		s.log.Error("failed to save user", zap.Error(err))
		return 0, err
	}

	s.log.Info("saved employee", zap.Int("id", id))
	return id, nil
}
