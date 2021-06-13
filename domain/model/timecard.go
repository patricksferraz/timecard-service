package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Timecard struct {
	Base          `valid:"required"`
	Status        TimecardStatus `json:"status" gorm:"column:status;not null" valid:"timecardStatus"`
	StartDate     time.Time      `json:"start_date" gorm:"column:start_date;not null" valid:"required"`
	EndDate       time.Time      `json:"end_date" gorm:"column:end_date;not null" valid:"required"`
	WorkedHours   time.Duration  `json:"worked_hours,omitempty" gorm:"column:worked_hours" valid:"-"`
	OvertimeHours time.Duration  `json:"overtime_hours,omitempty" gorm:"column:overtime_hours" valid:"-"`
	AbsenceHours  time.Duration  `json:"absence_hours,omitempty" gorm:"column:absence_hours" valid:"-"`
	TimeRecords   []*TimeRecord  `json:"time_records,omitempty" gorm:"column:time_records;foreignKey:TimecardID" valid:"-"`
	Company       *Company       `valid:"-"`
	CompanyID     string         `json:"company_id" gorm:"column:company_id;not null" valid:"uuid"`
	Employee      *Employee      `valid:"-"`
	EmployeeID    string         `json:"employee_id" gorm:"column:employee_id;type:uuid;not null" valid:"uuid"`
	Auditor       *Employee      `valid:"-"`
	AuditedBy     string         `json:"audited_by,omitempty" gorm:"column:audited_by;type:uuid" valid:"-"`
	Approver      *Employee      `valid:"-"`
	ApprovedBy    string         `json:"approved_by,omitempty" gorm:"column:approved_by;type:uuid" valid:"-"`
}

func NewTimecard(startDate time.Time, endDate time.Time, company *Company, employee *Employee) (*Timecard, error) {

	timecard := Timecard{
		Status:     TIMECARD_OPEN,
		StartDate:  startDate,
		EndDate:    endDate,
		Company:    company,
		CompanyID:  company.ID,
		Employee:   employee,
		EmployeeID: employee.ID,
	}

	timecard.ID = uuid.NewV4().String()
	timecard.CreatedAt = time.Now()

	err := timecard.isValid()
	if err != nil {
		return nil, err
	}

	return &timecard, nil
}

func (t *Timecard) isValid() error {
	_, err := govalidator.ValidateStruct(t)
	return err
}

func (t *Timecard) Calculate() {
}

func (t *Timecard) WaitApproval(auditor *Employee) error {
	if t.Status == TIMECARD_APPROVED {
		return errors.New("the approved timecard cannot be change")
	}

	if t.Status == TIMECARD_AWAITING_APPROVAL {
		return errors.New("the timecard is already awaiting approval")
	}

	t.Calculate()
	t.Status = TIMECARD_AWAITING_APPROVAL
	t.Auditor = auditor
	t.AuditedBy = auditor.ID
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Timecard) Approve(approver *Employee) error {
	if t.Status == TIMECARD_APPROVED {
		return errors.New("the approved timecard cannot be change")
	}

	t.Status = TIMECARD_APPROVED
	t.Approver = approver
	t.ApprovedBy = approver.ID
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Timecard) AddTimeRecord(timeRecord *TimeRecord) error {
	if t.Status != TIMECARD_OPEN {
		return errors.New("the timecard can only be changed when opened")
	}

	if timeRecord.EmployeeID != t.EmployeeID {
		return errors.New("the card and time record must belong to the same employee")
	}

	if timeRecord.Status != TIME_RECORD_APPROVED {
		return errors.New("the time record must be approved")
	}

	if timeRecord.Time.Before(t.StartDate) || timeRecord.Time.After(t.EndDate) {
		text := fmt.Sprintf("the time record must belong to the range %s to %s", t.StartDate, t.EndDate)
		return errors.New(text)
	}

	t.TimeRecords = append(t.TimeRecords, timeRecord)
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}
