package sales

type Sales struct {
	Position string
	Salary   int
	Address  string
}

func (s *Sales) GetPosition() string {
	return s.Position
}

func (s *Sales) SetPosition(position string) {
	s.Position = position
}

func (s *Sales) GetSalary() int {
	return s.Salary
}

func (s *Sales) SetSalary(salary int) {
	s.Salary = salary
}

func (s *Sales) GetAddress() string {
	return s.Address
}

func (s *Sales) SetAddress(address string) {
	s.Address = address
}
