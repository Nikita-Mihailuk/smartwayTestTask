package postgres

import (
	"context"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/model"
)

func (s *Storage) GetEmployeesByCompanyID(ctx context.Context, companyID int) ([]model.Employee, error) {
	// проверка есть ли компания с данным id
	var exists bool
	err := s.db.QueryRow(ctx, queryFindCompanyExist, companyID).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrCompanyNotFound
	}

	// если есть, то ищем сотрудников
	rows, err := s.db.Query(ctx, queryGetEmployeesByCompanyID, companyID)
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
