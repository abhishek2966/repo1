package user

// Employee struct
type Employee struct {
	EmployeeId  int
	Name, Email string
}

func (e *Employee) ChangeName(newName string) {
	e.Name = newName
}
