package schema

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type CompanyEmployeeEvent struct {
	Event      `json:",inline" valid:"required"`
	CompanyID  string `json:"company_id,omitempty" valid:"uuid"`
	EmployeeID string `json:"employee_id,omitempty" valid:"uuid"`
}

func NewCompanyEmployeeEvent() *CompanyEmployeeEvent {
	return &CompanyEmployeeEvent{}
}
