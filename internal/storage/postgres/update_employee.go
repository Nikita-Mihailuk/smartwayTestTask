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

func (s *Storage) UpdateEmployee(ctx context.Context, update dto.UpdateEmployee) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// получаем passport_id и при этом проверяем существование сотрудника
	var passportID int
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
		err = tx.QueryRow(ctx, queryFindCompanyExist, update.CompanyID).Scan(&exists)
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
		err = tx.QueryRow(ctx, queryFindDepartmentExist, update.DepartmentID).Scan(&exists)
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
		err = tx.QueryRow(ctx, queryDuplicatePassport, newType, newNumber, passportID).Scan(&exists)
		if err != nil {
			return err
		}
		if exists {
			return ErrPassportExist
		}

		// обновляем паспорт
		_, err = tx.Exec(ctx, queryUpdatePassport, newType, newNumber, passportID)
		if err != nil {
			return err
		}
	}

	// обновляем сотрудника
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
