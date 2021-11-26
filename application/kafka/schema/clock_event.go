package schema

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ClockEvent struct {
	Event `json:",inline" valid:"required"`
	Clock *Clock `json:"clock,omitempty" valid:"required"`
}

func NewClockEvent() *ClockEvent {
	return &ClockEvent{}
}

func (e *ClockEvent) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *ClockEvent) ParseJson(data []byte) error {
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

type DeleteClockEvent struct {
	Event       `json:",inline" valid:"required"`
	CompanyID   string `json:"company_id" valid:"uuid"`
	WorkScaleID string `json:"work_scale_id" valid:"uuid"`
	ClockID     string `json:"clock_id" valid:"uuid"`
}

func NewDeleteClockEvent() *DeleteClockEvent {
	return &DeleteClockEvent{}
}

func (e *DeleteClockEvent) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *DeleteClockEvent) ParseJson(data []byte) error {
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
