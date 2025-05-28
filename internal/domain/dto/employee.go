package dto

type UpdateEmployee struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Surname        string `json:"surname,omitempty"`
	Phone          string `json:"phone,omitempty"`
	CompanyID      int    `json:"company_id,omitempty"`
	PassportType   string `json:"passport_type,omitempty"`
	PassportNumber string `json:"passport_number,omitempty"`
	DepartmentID   int    `json:"department_id,omitempty"`
}
