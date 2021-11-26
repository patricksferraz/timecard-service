package schema

type Employee struct {
	Base `json:",inline" valid:"required"`
}

func NewEmployee() *Employee {
	return &Employee{}
}
