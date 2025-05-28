package model

type Employee struct {
	ID         int        `json:"id,omitempty"`
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Phone      string     `json:"phone"`
	CompanyID  int        `json:"company_id"`
	Passport   Passport   `json:"passport"`
	Department Department `json:"department"`
}

type Passport struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Department struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
