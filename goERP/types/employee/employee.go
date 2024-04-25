package employeeType

type Employee struct {
	Id int64 `json:"id"`
	AddEmployee
}

type Employees []Employee
