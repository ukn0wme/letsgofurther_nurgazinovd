package hr

type HR struct {
	Position string
	Salary   int
	Address  string
}

func (h *HR) GetPosition() string {
	return h.Position
}

func (h *HR) SetPosition(position string) {
	h.Position = position
}

func (h *HR) GetSalary() int {
	return h.Salary
}

func (h *HR) SetSalary(salary int) {
	h.Salary = salary
}

func (h *HR) GetAddress() string {
	return h.Address
}

func (h *HR) SetAddress(address string) {
	h.Address = address
}
