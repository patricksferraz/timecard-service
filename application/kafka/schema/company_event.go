package schema

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type CompanyEvent struct {
	Event   `json:",inline" valid:"required"`
	Company *Company `json:"company,omitempty" valid:"required"`
}

func NewCompanyEvent() *CompanyEvent {
	return &CompanyEvent{}
}
