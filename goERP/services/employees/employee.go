package employeeService

import employeeRepository "goERP/repositories/employee"

type EmployeeService struct {
	addEmployeeRepository    employeeRepository.AddEmployee
	updateEmployeeRepository employeeRepository.UpdateEmployee
}

func MakeAddEmployeeService(employeeRepository employeeRepository.AddEmployee) *EmployeeService {
	return &EmployeeService{
		addEmployeeRepository: employeeRepository,
	}
}

func MakeUpdateEmployeeService(employeeRepository employeeRepository.UpdateEmployee) *EmployeeService {
	return &EmployeeService{
		updateEmployeeRepository: employeeRepository,
	}
}
