package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

func (s *Storage) DeleteEmployee(ctx context.Context, id int) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// получаем id паспорта для удаления, это сразу же проверяет есть ли пользователь с таким id
	var passportID int
	err = tx.QueryRow(ctx, queryFindPassportID, id).Scan(&passportID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrEmployeeNotFound
		}
		return err
	}

	// сначала удаляем сотрудника
	_, err = tx.Exec(ctx, queryDeleteEmployee, id)
	if err != nil {
		return err
	}

	// удаляем паспорт
	_, err = tx.Exec(ctx, queryDeletePassport, passportID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
