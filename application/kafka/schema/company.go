package schema

type Company struct {
	Base `json:",inline" valid:"required"`
}

func NewCompany() *Company {
	return &Company{}
}
