package schema

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type WorkScaleEmployeeEvent struct {
	Event       `json:",inline" valid:"required"`
	CompanyID   string `json:"company_id,omitempty" valid:"uuid"`
	EmployeeID  string `json:"employee_id,omitempty" valid:"uuid"`
	WorkScaleID string `json:"work_scale_id,omitempty" valid:"uuid"`
}

func NewWorkScaleEmployeeEvent() *WorkScaleEmployeeEvent {
	return &WorkScaleEmployeeEvent{}
}

func (e *WorkScaleEmployeeEvent) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *WorkScaleEmployeeEvent) ParseJson(data []byte) error {
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
