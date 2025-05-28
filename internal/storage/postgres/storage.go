package postgres

import (
	"context"
	"errors"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/dto"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/domain/model"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *Storage) SaveEmployee(ctx context.Context, employee model.Employee) (int, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	// проверка есть ли компания с данным id
	var exists bool
	queryFindCompany := `SELECT EXISTS(SELECT 1 FROM companies WHERE id = $1)`
	err = tx.QueryRow(ctx, queryFindCompany, employee.CompanyID).Scan(&exists)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, ErrCompanyNotFound
	}

	// проверяем есть такой паспорт, если нет то вставляем новый и возвращаем id записи
	queryFindPassport := ` SELECT EXISTS(SELECT 1 FROM passports WHERE number = $1 AND type = $2)`

	err = tx.QueryRow(ctx,
		queryFindPassport,
		employee.Passport.Number,
		employee.Passport.Type).
		Scan(&exists)

	if err != nil {
		return 0, err
	}
	if exists {
		return 0, ErrPassportExist
	}

	var passportID int
	queryInsertPassport := `
		INSERT INTO passports (type, number)
		VALUES ($1, $2)
		RETURNING id
	`
	err = tx.QueryRow(ctx,
		queryInsertPassport,
		employee.Passport.Type,
		employee.Passport.Number).
		Scan(&passportID)

	if err != nil {
		return 0, err
	}

	// проверка есть ли данный отдел, если нет то создаем его
	var departmentID int
	queryFindDepartment := `
		SELECT id FROM departments
		WHERE name = $1 AND phone = $2 AND company_id = $3
	`
	err = tx.QueryRow(ctx,
		queryFindDepartment,
		employee.Department.Name,
		employee.Department.Phone,
		employee.CompanyID).
		Scan(&departmentID)

	if err != nil {
		queryInsertDepartment := `
			INSERT INTO departments (name, phone, company_id)
			VALUES ($1, $2, $3)
			RETURNING id
		`
		err = tx.QueryRow(ctx,
			queryInsertDepartment,
			employee.Department.Name,
			employee.Department.Phone,
			employee.CompanyID).
			Scan(&departmentID)

		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
				return 0, ErrDepartmentExist
			}
			return 0, err
		}
	}

	// вставляем сотрудника
	var employeeID int
	queryEmployee := `
		INSERT INTO employees (name, surname, phone, company_id, passport_id, department_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	err = tx.QueryRow(ctx, queryEmployee,
		employee.Name,
		employee.Surname,
		employee.Phone,
		employee.CompanyID,
		passportID,
		departmentID).
		Scan(&employeeID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return 0, ErrEmployeeExist
		}
		return 0, err
	}

	if err = tx.Commit(ctx); err != nil {
		return 0, err
	}

	return employeeID, nil
}

func (s *Storage) GetEmployeesByCompanyID(ctx context.Context, companyID int) ([]model.Employee, error) {
	// проверка есть ли компания с данным id
	var exists bool
	queryFindCompany := `SELECT EXISTS(SELECT 1 FROM companies WHERE id = $1)`
	err := s.db.QueryRow(ctx, queryFindCompany, companyID).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrCompanyNotFound
	}

	// если есть, то ищем сотрудников
	queryGetEmployeesByCompanyID := `
		SELECT employees.name, employees.surname, employees.phone, employees.company_id, passports.type, passports.number, departments.name, departments.phone
    	FROM employees 
    	JOIN passports ON employees.passport_id = passports.id
    	JOIN departments ON employees.department_id = departments.id
    	WHERE employees.company_id=$1
	`

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

func (s *Storage) GetEmployeeByDepartmentID(ctx context.Context, departmentID, companyID int) ([]model.Employee, error) {
	// проверка есть ли такой отдел в данной компании
	var exists bool
	queryFindDepartment := `SELECT EXISTS(SELECT 1 FROM departments WHERE id = $1 AND company_id = $2)`
	err := s.db.QueryRow(ctx, queryFindDepartment, departmentID, companyID).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrDepartmentNotFound
	}

	// если есть, то ищем сотрудников
	queryGetEmployeesByCompanyID := `
		SELECT employees.name, employees.surname, employees.phone, employees.company_id, passports.type, passports.number, departments.name, departments.phone
    	FROM employees 
    	JOIN passports ON employees.passport_id = passports.id
    	JOIN departments ON employees.department_id = departments.id
    	WHERE employees.company_id=$1 AND employees.department_id=$2
	`

	rows, err := s.db.Query(ctx, queryGetEmployeesByCompanyID, companyID, departmentID)
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

func (s *Storage) DeleteEmployee(ctx context.Context, id int) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// получаем id паспорта для удаления, это сразу же проверяет есть ли пользователь с таким id
	var passportID int
	queryFindPassportID := `SELECT passport_id FROM employees WHERE id = $1`
	err = tx.QueryRow(ctx, queryFindPassportID, id).Scan(&passportID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrEmployeeNotFound
		}
		return err
	}

	// сначала удаляем сотрудника
	queryDeleteEmployee := `DELETE FROM employees WHERE id = $1`
	_, err = tx.Exec(ctx, queryDeleteEmployee, id)
	if err != nil {
		return err
	}

	// удаляем паспорт
	queryDeletePassport := `DELETE FROM passports WHERE id = $1`
	_, err = tx.Exec(ctx, queryDeletePassport, passportID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *Storage) UpdateEmployee(ctx context.Context, update dto.UpdateEmployee) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// получаем passport_id и при этом проверяем существование сотрудника
	var passportID int
	queryFindPassportID := `SELECT passport_id FROM employees WHERE id = $1`
	err = tx.QueryRow(ctx, queryFindPassportID, update.ID).Scan(&passportID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrEmployeeNotFound
		}
		return err
	}

	// получаем текущие значения полей сотрудника
	var current model.Employee
	var departmentID int
	queryFindEmployee := `
		SELECT employees.name, employees.surname, employees.phone, employees.company_id, passports.type, passports.number, departments.name, departments.phone, employees.department_id
		FROM employees
		JOIN passports ON employees.passport_id = passports.id
		JOIN departments ON employees.department_id = departments.id
		WHERE employees.id = $1
	`
	err = tx.QueryRow(ctx, queryFindEmployee, update.ID).Scan(
		&current.Name,
		&current.Surname,
		&current.Phone,
		&current.CompanyID,
		&current.Passport.Type,
		&current.Passport.Number,
		&current.Department.Name,
		&current.Department.Phone,
		&departmentID,
	)
	if err != nil {
		return err
	}

	// подставляем новые значения, если они были переданы
	if update.Name != "" {
		current.Name = update.Name
	}
	if update.Surname != "" {
		current.Surname = update.Surname
	}
	if update.Phone != "" {
		current.Phone = update.Phone
	}
	if update.CompanyID != 0 {
		// проверка есть ли компания с данным id
		var exists bool
		queryFindCompany := `SELECT EXISTS(SELECT 1 FROM companies WHERE id = $1)`
		err = tx.QueryRow(ctx, queryFindCompany, update.CompanyID).Scan(&exists)
		if err != nil {
			return err
		}
		if !exists {
			return ErrCompanyNotFound
		}
		current.CompanyID = update.CompanyID
	}
	if update.DepartmentID != 0 {
		// проверка есть ли такой отдел
		var exists bool
		queryFindDepartment := `SELECT EXISTS(SELECT 1 FROM departments WHERE id = $1)`
		err = tx.QueryRow(ctx, queryFindDepartment, update.DepartmentID).Scan(&exists)
		if err != nil {
			return err
		}
		if !exists {
			return ErrDepartmentNotFound
		}
		departmentID = update.DepartmentID
	}

	if update.PassportType != "" || update.PassportNumber != "" {

		// для случая когда передан либо только тип, либо только номер
		newType := current.Passport.Type
		if update.PassportType != "" {
			newType = update.PassportType
		}

		newNumber := current.Passport.Number
		if update.PassportNumber != "" {
			newNumber = update.PassportNumber
		}

		// проверка уникальности паспорта
		var exists bool
		queryDublicatePassport := `
			SELECT EXISTS(
				SELECT 1 FROM passports
				WHERE type = $1 AND number = $2 AND id <> $3
			)
		`
		err = tx.QueryRow(ctx, queryDublicatePassport, newType, newNumber, passportID).Scan(&exists)
		if err != nil {
			return err
		}
		if exists {
			return ErrPassportExist
		}

		// обновляем паспорт
		queryUpdatePassport := `
			UPDATE passports SET type = $1, number = $2
			WHERE id = $3
		`
		_, err = tx.Exec(ctx, queryUpdatePassport, newType, newNumber, passportID)
		if err != nil {
			return err
		}
	}

	// обновляем сотрудника
	queryUpdateEmployee := `
		UPDATE employees
		SET name = $1, surname = $2, phone = $3, company_id = $4, department_id = $5
		WHERE id = $6
	`
	_, err = tx.Exec(ctx,
		queryUpdateEmployee,
		current.Name,
		current.Surname,
		current.Phone,
		current.CompanyID,
		departmentID,
		update.ID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return ErrEmployeeExist
		}
		return err
	}

	return tx.Commit(ctx)
}
