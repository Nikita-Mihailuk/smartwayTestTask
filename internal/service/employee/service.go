package employee

import (
	"context"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/model"
	"go.uber.org/zap"
)

type EmployeeService struct {
	log              *zap.Logger
	employeeSaver    EmployeeSaver
	employeeProvider EmployeeProvider
	employeeDeleter  EmployeeDeleter
	employeeUpdater  EmployeeUpdater
}

func NewEmployeeService(
	log *zap.Logger,
	employeeSaver EmployeeSaver,
	employeeProvider EmployeeProvider,
	employeeDeleter EmployeeDeleter,
	employeeUpdater EmployeeUpdater) *EmployeeService {

	return &EmployeeService{
		log:              log,
		employeeSaver:    employeeSaver,
		employeeProvider: employeeProvider,
		employeeDeleter:  employeeDeleter,
		employeeUpdater:  employeeUpdater,
	}
}

type EmployeeSaver interface {
	SaveEmployee(ctx context.Context, employee model.Employee) (int, error)
}

type EmployeeProvider interface {
	GetEmployeesByCompanyID(ctx context.Context, companyID int) ([]model.Employee, error)
	GetEmployeeByDepartmentID(ctx context.Context, departmentID, companyID int) ([]model.Employee, error)
}

type EmployeeDeleter interface {
	DeleteEmployee(ctx context.Context, id int) error
}

type EmployeeUpdater interface {
	UpdateEmployee(ctx context.Context, employee model.Employee) error
}
