package customerService

import (
	"github.com/go-playground/validator/v10"
	"goERP/services/dataProcessingUtils"
	"goERP/types/customer"
	"goERP/types/restResponses"
	inputValidations "goERP/validations"
	"net/http"
)

func (cs *CustomerService) Update(customer customer.AddCustomer, id int64) (*customer.Customer, *restResponses.Error) {
	errOnValidate := inputValidations.Validate.Struct(customer)
	if errOnValidate != nil {
		for _, err := range errOnValidate.(validator.ValidationErrors) {
			return nil, &restResponses.Error{
				Status:  http.StatusBadRequest,
				Field:   err.Field(),
				Message: err.Error(),
			}
		}
	}

	dataProcessingUtils.RemoveNonInteger(
		&customer.Cellphone,
		&customer.Phone,
		&customer.CpfCnpj,
	)

	updatedCustomer, err := cs.updateCustomerRepository.Update(customer, id)
	if err != nil {
		return nil, &restResponses.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &updatedCustomer, nil
}
