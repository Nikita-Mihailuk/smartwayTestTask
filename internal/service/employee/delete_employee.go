package employee

import (
	"context"
	"errors"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/storage/postgres"
	"go.uber.org/zap"
)

func (s *EmployeeService) DropEmployee(ctx context.Context, id int) error {
	err := s.employeeDeleter.DeleteEmployee(ctx, id)
	if err != nil {
		if errors.Is(err, postgres.ErrEmployeeNotFound) {
			s.log.Error("employee not found", zap.Error(err))
			return ErrEmployeeNotFoundByID
		}
		s.log.Error("failed to delete employee", zap.Error(err))
		return err
	}

	s.log.Info("deleted employee", zap.Int("id", id))
	return nil
}
