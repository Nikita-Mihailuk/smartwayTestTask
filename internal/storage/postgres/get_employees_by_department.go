package postgres

import (
	"context"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/model"
)

func (s *Storage) GetEmployeeByDepartmentID(ctx context.Context, departmentID, companyID int) ([]model.Employee, error) {
	// проверка есть ли такой отдел в данной компании
	var exists bool
	err := s.db.QueryRow(ctx, queryFindDepartmentByCID, departmentID, companyID).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrDepartmentNotFound
	}

	// если есть, то ищем сотрудников
	rows, err := s.db.Query(ctx, queryGetEmployeesByDepartmentID, companyID, departmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		err = rows.Scan(&employee.Name,
			&employee.Surname,
			&employee.Phone,
			&employee.CompanyID,
			&employee.Passport.Type,
			&employee.Passport.Number,
			&employee.Department.Name,
			&employee.Department.Phone)

		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	if len(employees) == 0 {
		return nil, ErrEmployeesNotFound
	}

	return employees, nil
}
