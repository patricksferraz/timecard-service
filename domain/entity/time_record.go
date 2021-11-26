package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type TimeRecord struct {
	Base       `bson:",inline" valid:"-"`
	Time       time.Time        `json:"time,omitempty" gorm:"column:time;not null;unique_index:idx_employee_company_time" valid:"required"`
	Status     TimeRecordStatus `json:"status" gorm:"column:status;not null" valid:"timeRecordStatus"`
	TzOffset   int              `json:"tz_offset" valid:"int,optional"`
	EmployeeID *string          `json:"employee_id,omitempty" gorm:"column:employee_id;type:uuid;not null;unique_index:idx_employee_company_time" bson:"employee_id" valid:"uuid"`
	Employee   *Employee        `json:"-" valid:"-"`
	CompanyID  *string          `json:"company_id,omitempty" gorm:"column:company_id;type:uuid;not null;unique_index:idx_employee_company_time" bson:"company_id" valid:"uuid"`
	Company    *Company         `json:"-" valid:"-"`
}

func NewTimeRecord(id string, _time time.Time, tzOffset int, status int, employee *Employee, company *Company) (*TimeRecord, error) {
	timeRecord := TimeRecord{
		Time:       _time,
		Status:     TimeRecordStatus(status),
		TzOffset:   tzOffset,
		EmployeeID: &employee.ID,
		Employee:   employee,
		CompanyID:  &company.ID,
		Company:    company,
	}

	timeRecord.ID = id
	timeRecord.CreatedAt = time.Now()

	err := timeRecord.isValid()
	if err != nil {
		return nil, err
	}

	return &timeRecord, nil
}

func (t *TimeRecord) isValid() error {
	_, err := govalidator.ValidateStruct(t)
	return err
}

func (t *TimeRecord) Approve() error {
	t.Status = TIME_RECORD_APPROVED
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *TimeRecord) Refuse() error {
	t.Status = TIME_RECORD_REFUSED
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}
