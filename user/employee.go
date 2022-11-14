package user

// Employee struct
type Employee struct {
	EmployeeId  int
	Name, email string
}

func (e *Employee) ChangeName(newName string) {
	e.Name = newName
}

func (e *Employee) SetEmail(email string) {

	e.email = email
}
