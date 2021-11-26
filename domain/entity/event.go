package entity

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Event struct {
	Base     `json:",inline" valid:"required"`
	Status   EventStatus `json:"status" gorm:"column:status;not null" valid:"eventStatus"`
	Resume   *string     `json:"resume,omitempty" gorm:"column:value;type:varchar(100)" valid:"-"`
	Attempts int         `json:"attemps,omitempty" gorm:"column:attempts" valid:"-"`
}

func NewEvent(id, resume string) (*Event, error) {
	company := Event{
		Resume:   &resume,
		Status:   EVENT_PENDING,
		Attempts: 1,
	}
	company.ID = id
	company.CreatedAt = time.Now()

	if err := company.isValid(); err != nil {
		return nil, err
	}

	return &company, nil
}

func (e *Event) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *Event) AddAttempt() error {

	if e.Status == EVENT_COMPLETED {
		return errors.New("event is completed")
	}

	if e.Status == EVENT_FAILED {
		return errors.New("event is failed")
	}

	if e.Attempts >= 10 {
		e.Status = EVENT_FAILED
		return errors.New("event has reached the maximum number of attempts")
	}

	e.Attempts++

	return nil
}

func (e *Event) Complete() error {

	if e.Status == EVENT_COMPLETED {
		return errors.New("the event has already been completed")
	}

	if e.Status == EVENT_FAILED {
		return errors.New("the failed event cannot be completed")
	}

	e.Status = EVENT_COMPLETED
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Event) Fail() error {

	if e.Status == EVENT_COMPLETED {
		return errors.New("the completed event cannot be failed")
	}

	if e.Status == EVENT_FAILED {
		return errors.New("the event has already been failed")
	}

	e.Status = EVENT_FAILED
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}
