package customer

import cityType "goERP/types/location"

type AddCustomer struct {
	Name           string                `json:"name" validate:"required"`
	CompanyName    string                `json:"companyName"`
	CpfCnpj        string                `json:"cpf_cnpj"`
	Phone          string                `json:"phone"`
	Cellphone      string                `json:"cellphone"`
	Email          string                `json:"email" validate:"omitempty,email"`
	Obs            string                `json:"obs"`
	Location       cityType.FullLocation `json:"location"`
	PriceTableCode int                   `json:"priceTableId"`
}
