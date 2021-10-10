package schema

type Employee struct {
	Base `json:",inline" valid:"required"`
}

func NewEmployee(id, pis string) *Employee {
	return &Employee{}
}
