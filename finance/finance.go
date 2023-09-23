package finance

type Finance struct {
	Position string
	Salary   int
	Address  string
}

func (f *Finance) GetPosition() string {
	return f.Position
}

func (f *Finance) SetPosition(position string) {
	f.Position = position
}

func (f *Finance) GetSalary() int {
	return f.Salary
}

func (f *Finance) SetSalary(salary int) {
	f.Salary = salary
}

func (f *Finance) GetAddress() string {
	return f.Address
}

func (f *Finance) SetAddress(address string) {
	f.Address = address
}
