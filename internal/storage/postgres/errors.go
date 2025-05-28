package postgres

import "errors"

var (
	ErrEmployeeExist      = errors.New("employee already exists")
	ErrDepartmentExist    = errors.New("department already exists")
	ErrPassportExist      = errors.New("passport already exists")
	ErrEmployeesNotFound  = errors.New("employees not found")
	ErrCompanyNotFound    = errors.New("company not found")
	ErrDepartmentNotFound = errors.New("department not found")
	ErrPassportNotFound   = errors.New("passport not found")
)
