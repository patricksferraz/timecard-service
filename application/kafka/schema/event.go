package schema

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
)

type Event struct {
	ID string `json:"id,omitempty" valid:"uuid"`
}

func NewEvent(id, pis string) *Event {
	return &Event{}
}

func (e *Event) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *Event) ParseJson(data []byte) error {
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
