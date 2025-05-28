package employee

import "errors"

var (
	ErrEmployeeExist                 = errors.New("employee already exists")
	ErrDepartmentExist               = errors.New("department already exists")
	ErrPassportExist                 = errors.New("passport already exists")
	ErrEmployeesNotFoundByCompany    = errors.New("employees not found by company")
	ErrEmployeesNotFoundByDepartment = errors.New("employees not found by department")
	ErrInvalidCompany                = errors.New("invalid company")
	ErrInvalidDepartment             = errors.New("invalid department")
	ErrInvalidPassport               = errors.New("invalid passport")
)
