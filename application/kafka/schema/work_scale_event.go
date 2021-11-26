package schema

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
)

type WorkScaleEvent struct {
	Event     `json:",inline" valid:"required"`
	WorkScale *WorkScale `json:"work_scale,omitempty" valid:"-"`
}

func NewWorkScaleEvent() *WorkScaleEvent {
	return &WorkScaleEvent{}
}

func (e *WorkScaleEvent) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *WorkScaleEvent) ParseJson(data []byte) error {
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
