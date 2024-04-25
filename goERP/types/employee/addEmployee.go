package employeeType

import cityType "goERP/types/location"

type SalaryType string

const (
	Monthly SalaryType = "M"
	Hourly  SalaryType = "H"
)

type AddEmployee struct {
	Name       string                `json:"name" validate:"required"`
	CpfCnpj    string                `json:"cpf_cnpj"`
	Phone      string                `json:"phone"`
	Cellphone  string                `json:"cellphone"`
	Email      string                `json:"email" validate:"omitempty,email"`
	Obs        string                `json:"obs"`
	Location   cityType.FullLocation `json:"location"`
	SalaryType SalaryType            `json:"salary_type" validate:"required,alpha,oneof=H M"`
	Salary     float32               `json:"salary" validate:"required,gt=0,numeric"`
}
