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
	Topic    *string     `json:"topic,omitempty" gorm:"column:topic;type:varchar(100)" valid:"-"`
	Attempts *int        `json:"attemps,omitempty" gorm:"column:attempts" valid:"-"`
}

func NewEvent(id, topic string) (*Event, error) {
	company := Event{
		Topic:  &topic,
		Status: EVENT_PENDING,
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

func (e *Event) AddAttempt() {
	if e.Attempts == nil {
		e.Attempts = new(int)
	}

	*e.Attempts++
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

func (e *Event) IsCompleted() bool {
	return e.Status == EVENT_COMPLETED
}

func (e *Event) IsFailed() bool {
	return e.Status == EVENT_FAILED
}
