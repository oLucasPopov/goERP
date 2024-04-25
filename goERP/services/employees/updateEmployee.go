package employeeService

import (
	"github.com/go-playground/validator/v10"
	"goERP/services/dataProcessingUtils"
	employeeType "goERP/types/employee"
	"goERP/types/restResponses"
	inputValidations "goERP/validations"
)

func (es *EmployeeService) Update(employee employeeType.AddEmployee, id int64) (*employeeType.Employee, *restResponses.Error) {
	errOnValidate := inputValidations.Validate.Struct(employee)
	if errOnValidate != nil {
		for _, err := range errOnValidate.(validator.ValidationErrors) {
			return nil, &restResponses.Error{
				Field:   err.Field(),
				Message: err.Error(),
			}
		}
	}

	dataProcessingUtils.RemoveNonInteger(
		&employee.Cellphone,
		&employee.Phone,
		&employee.CpfCnpj,
	)

	newEmployee, err := es.updateEmployeeRepository.Update(employee, id)
	if err != nil {
		return nil, &restResponses.Error{
			Message: err.Error(),
		}
	}

	return &newEmployee, nil
}
