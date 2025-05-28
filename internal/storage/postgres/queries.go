package postgres

const (
	queryFindCompanyExist        = `SELECT EXISTS(SELECT 1 FROM companies WHERE id = $1)`
	queryFindDepartment          = `SELECT id FROM departments WHERE phone = $1 AND name = $2 AND company_id = $3`
	queryFindDepartmentExist     = `SELECT EXISTS(SELECT 1 FROM departments WHERE id = $1)`
	queryFindDepartmentByCID     = `SELECT EXISTS(SELECT 1 FROM departments WHERE id = $1 AND company_id = $2)`
	queryFindPassport            = `SELECT EXISTS(SELECT 1 FROM passports WHERE number = $1 AND type = $2)`
	queryInsertPassport          = `INSERT INTO passports (type, number) VALUES ($1, $2) RETURNING id`
	queryInsertEmployee          = `INSERT INTO employees (name, surname, phone, company_id, passport_id, department_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	queryGetEmployeesByCompanyID = `
		SELECT employees.name, employees.surname, employees.phone, employees.company_id,
		passports.type, passports.number, departments.name, departments.phone
		FROM employees 
		JOIN passports ON employees.passport_id = passports.id
		JOIN departments ON employees.department_id = departments.id
		WHERE employees.company_id=$1
	`
	queryGetEmployeesByDepartmentID = `
		SELECT employees.name, employees.surname, employees.phone, employees.company_id,
		passports.type, passports.number, departments.name, departments.phone
		FROM employees 
		JOIN passports ON employees.passport_id = passports.id
		JOIN departments ON employees.department_id = departments.id
		WHERE employees.company_id=$1 AND employees.department_id=$2
	`
	queryDeleteEmployee = `DELETE FROM employees WHERE id = $1`
	queryDeletePassport = `DELETE FROM passports WHERE id = $1`
	queryFindPassportID = `SELECT passport_id FROM employees WHERE id = $1`
	queryFindEmployee   = `
		SELECT employees.name, employees.surname, employees.phone, employees.company_id,
		passports.type, passports.number, departments.name, departments.phone, employees.department_id
		FROM employees
		JOIN passports ON employees.passport_id = passports.id
		JOIN departments ON employees.department_id = departments.id
		WHERE employees.id = $1
	`
	queryUpdatePassport    = `UPDATE passports SET type = $1, number = $2 WHERE id = $3`
	queryUpdateEmployee    = `UPDATE employees SET name = $1, surname = $2, phone = $3, company_id = $4, department_id = $5 WHERE id = $6`
	queryDuplicatePassport = `SELECT EXISTS(SELECT 1 FROM passports WHERE type = $1 AND number = $2 AND id <> $3)`
)
