package magazine

type Subsriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

type Employee struct {
	Name     string
	Salary   float64
	HomeAddr Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
