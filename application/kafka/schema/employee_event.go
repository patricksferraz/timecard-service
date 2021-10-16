package schema

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type EmployeeEvent struct {
	Event    `json:",inline" valid:"required"`
	Employee *Employee `json:"employee,omitempty" valid:"-"`
}

func NewEmployeeEvent(id, pis string) *EmployeeEvent {
	return &EmployeeEvent{}
}
