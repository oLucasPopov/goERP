package main

import (
	"fmt"
	"github.com/gorilla/mux"
	migrations "goERP/config/migration"
	"goERP/controllers/customers"
	"goERP/controllers/employees"
	"goERP/controllers/locations"
	"goERP/repositories/city"
	customerRepository "goERP/repositories/customer"
	employeeRepository "goERP/repositories/employee"
	customerService "goERP/services/customers"
	employeeService "goERP/services/employees"
	locationService "goERP/services/locations"
	"log"
	"net/http"
	"os"
)

func init() {
	configLogFile()
	migrations.Migrate()
}

func configLogFile() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc(
		"/employee/new",
		employees.MakeAddEmployee(*employeeService.MakeAddEmployeeService(employeeRepository.AddEmployee{})).Handle,
	).Methods(http.MethodPost)

	r.HandleFunc(
		"/employee/update/{id}",
		employees.MakeUpdateEmployee(*employeeService.MakeUpdateEmployeeService(employeeRepository.UpdateEmployee{})).Handle,
	).Methods(http.MethodPut)

	r.HandleFunc(
		"/customer/new",
		customers.MakeAddCustomer(*customerService.MakeAddCustomerService(customerRepository.AddCustomer{})).Handle,
	).Methods(http.MethodPost)

	r.HandleFunc(
		"/customer/delete/{id}",
		customers.MakeDeleteCustomer(*customerService.MakeDeleteCustomerService(customerRepository.DeleteCustomer{})).Handle,
	).Methods(http.MethodDelete)

	r.HandleFunc(
		"/customer/update/{id}",
		customers.MakeUpdateCustomer(*customerService.MakeUpdateCustomerService(customerRepository.UpdateCustomer{})).Handle,
	).Methods(http.MethodPut)

	r.HandleFunc(
		"/customer/get/{id}",
		customers.MakeGetCustomer(*customerService.MakeGetCustomerService(customerRepository.GetCustomer{})).Handle,
	).Methods(http.MethodGet)

	r.HandleFunc(
		"/customer/list/{id}",
		customers.MakeListCustomer(*customerService.MakeListCustomerService(customerRepository.ListCustomer{})).Handle,
	).Methods(http.MethodGet)

	r.HandleFunc(
		"/cities/get/{ibge-id}",
		locations.MakeGetCities(*locationService.MakeGetCityService(city.GetCities{})).Handle,
	).Methods(http.MethodGet)

	serverAddr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	log.Println("running on ", serverAddr)

	srv := &http.Server{
		Handler: r,
		Addr:    serverAddr,
	}

	log.Fatalln(srv.ListenAndServe())
}
