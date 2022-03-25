package user

type Employee struct {
	Id          int
	Name, Email string
}

func (e *Employee) ChangeName(newName string) {
	e.Name = newName
}
