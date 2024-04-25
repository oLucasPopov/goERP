package customer

type Customer struct {
	Id int64 `json:"id"`
	AddCustomer
}

type Customers []Customer
