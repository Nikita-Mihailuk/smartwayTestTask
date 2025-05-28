package postgres

import (
	"context"
	"errors"
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
	err = tx.QueryRow(ctx, queryFindCompanyExist, employee.CompanyID).Scan(&exists)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, ErrCompanyNotFound
	}

	// проверяем есть такой паспорт, если нет то вставляем новый и возвращаем id записи
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
	err = tx.QueryRow(ctx,
		queryFindDepartment,
		employee.Department.Phone,
		employee.Department.Name,
		employee.CompanyID).
		Scan(&departmentID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, ErrDepartmentNotFound
		}
	}

	// вставляем сотрудника
	var employeeID int
	err = tx.QueryRow(ctx,
		queryInsertEmployee,
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
