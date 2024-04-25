package customerService

import customerRepository "goERP/repositories/customer"

type CustomerService struct {
	addCustomerRepository    customerRepository.AddCustomer
	deleteCustomerRepository customerRepository.DeleteCustomer
	updateCustomerRepository customerRepository.UpdateCustomer
	getCustomerRepository    customerRepository.GetCustomer
	listCustomerRepository   customerRepository.ListCustomer
}

func MakeAddCustomerService(customerRepository customerRepository.AddCustomer) *CustomerService {
	return &CustomerService{
		addCustomerRepository: customerRepository,
	}
}

func MakeDeleteCustomerService(customerRepository customerRepository.DeleteCustomer) *CustomerService {
	return &CustomerService{
		deleteCustomerRepository: customerRepository,
	}
}

func MakeUpdateCustomerService(customerRepository customerRepository.UpdateCustomer) *CustomerService {
	return &CustomerService{
		updateCustomerRepository: customerRepository,
	}
}

func MakeGetCustomerService(customerRepository customerRepository.GetCustomer) *CustomerService {
	return &CustomerService{
		getCustomerRepository: customerRepository,
	}
}

func MakeListCustomerService(customerRepository customerRepository.ListCustomer) *CustomerService {
	return &CustomerService{
		listCustomerRepository: customerRepository,
	}
}
