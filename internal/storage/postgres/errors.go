package postgres

import "errors"

var (
	ErrEmployeeExist      = errors.New("employee already exists")
	ErrPassportExist      = errors.New("passport already exists")
	ErrEmployeesNotFound  = errors.New("employees not found")
	ErrCompanyNotFound    = errors.New("company not found")
	ErrDepartmentNotFound = errors.New("department not found")
	ErrEmployeeNotFound   = errors.New("employee not found")
)
