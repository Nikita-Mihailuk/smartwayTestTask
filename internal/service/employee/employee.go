package employee

import (
	"context"
	"errors"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/dto"
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
