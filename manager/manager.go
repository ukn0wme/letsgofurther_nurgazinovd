package manager

type Manager struct {
	Position string
	Salary   int
	Address  string
}

func (m *Manager) GetPosition() string {
	return m.Position
}

func (m *Manager) SetPosition(position string) {
	m.Position = position
}

func (m *Manager) GetSalary() int {
	return m.Salary
}

func (m *Manager) SetSalary(salary int) {
	m.Salary = salary
}

func (m *Manager) GetAddress() string {
	return m.Address
}

func (m *Manager) SetAddress(address string) {
	m.Address = address
}
