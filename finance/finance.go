package finance

type Finance struct {
	position string
	salary   int
	address  string
}

func (f *Finance) GetPosition() string {
	return f.position
}

func (f *Finance) SetPosition(position string) {
	f.position = position
}

func (f *Finance) GetSalary() int {
	return f.salary
}

func (f *Finance) SetSalary(salary int) {
	f.salary = salary
}

func (f *Finance) GetAddress() string {
	return f.address
}

func (f *Finance) SetAddress(address string) {
	f.address = address
}
