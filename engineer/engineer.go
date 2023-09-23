package engineer

type Engineer struct {
	Position string
	Salary   int
	Address  string
}

func (e *Engineer) GetPosition() string {
	return e.Position
}

func (e *Engineer) SetPosition(position string) {
	e.Position = position
}

func (e *Engineer) GetSalary() int {
	return e.Salary
}

func (e *Engineer) SetSalary(salary int) {
	e.Salary = salary
}

func (e *Engineer) GetAddress() string {
	return e.Address
}

func (e *Engineer) SetAddress(address string) {
	e.Address = address
}
