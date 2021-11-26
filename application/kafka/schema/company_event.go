package schema

import (
	"encoding/json"

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

func (e *CompanyEvent) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *CompanyEvent) ParseJson(data []byte) error {
	err := json.Unmarshal(data, e)
	if err != nil {
		return err
	}

	err = e.isValid()
	if err != nil {
		return err
	}

	return nil
}
