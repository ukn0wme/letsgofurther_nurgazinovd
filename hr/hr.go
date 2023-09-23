package hr

type HR struct {
	position string
	salary   int
	address  string
}

func (h *HR) GetPosition() string {
	return h.position
}

func (h *HR) SetPosition(position string) {
	h.position = position
}

func (h *HR) GetSalary() int {
	return h.salary
}

func (h *HR) SetSalary(salary int) {
	h.salary = salary
}

func (h *HR) GetAddress() string {
	return h.address
}

func (h *HR) SetAddress(address string) {
	h.address = address
}
