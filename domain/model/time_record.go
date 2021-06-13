package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type TimeRecord struct {
	Base       `valid:"required"`
	Time       time.Time        `json:"time" gorm:"column:time;type:time;not null" valid:"required"`
	Status     TimeRecordStatus `valid:"timeRecordStatus"`
	EmployeeID string           `valid:"uuid"`
	Timecard   *Timecard        `valid:"-"`
	TimecardID string           `json:"timecard_id" gorm:"column:timecard_id;type:uuid;not null" valid:"-"`
}

func NewTimeRecord(id string, _time time.Time, status TimeRecordStatus, employeeID string, timecard *Timecard) (*TimeRecord, error) {

	timeRecord := TimeRecord{
		Time:       _time,
		Status:     status,
		EmployeeID: employeeID,
		Timecard:   timecard,
		TimecardID: timecard.ID,
	}

	timeRecord.ID = id
	timeRecord.CreatedAt = time.Now()

	err := timeRecord.isValid()
	if err != nil {
		return nil, err
	}

	return &timeRecord, nil
}

func (p *TimeRecord) isValid() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
