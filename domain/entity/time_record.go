package entity

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// type TimeRecord struct {
// 	Base       `bson:",inline" valid:"-"`
// 	Time       time.Time        `json:"time,omitempty" gorm:"column:time;not null;unique_index:idx_employee_company_time" bson:"time" valid:"required"`
// 	Status     TimeRecordStatus `json:"status" gorm:"column:status;not null" bson:"status" valid:"timeRecordStatus"`
// 	EmployeeID *string          `json:"employee_id,omitempty" gorm:"column:employee_id;type:uuid;not null;unique_index:idx_employee_company_time" bson:"employee_id" valid:"uuid"`
// 	Employee   *Employee        `json:"-" valid:"-"`
// 	CompanyID  *string          `json:"company_id,omitempty" gorm:"column:company_id;type:uuid;not null;unique_index:idx_employee_company_time" bson:"company_id" valid:"uuid"`
// 	Company    *Company         `json:"-" valid:"-"`
// }

// func NewTimeRecord(id string, _time time.Time, status TimeRecordStatus, employee *Employee, company *Company) (*TimeRecord, error) {
// 	timeRecord := TimeRecord{
// 		Time:       _time,
// 		Status:     status,
// 		EmployeeID: &employee.ID,
// 		Employee:   employee,
// 		CompanyID:  &company.ID,
// 		Company:    company,
// 	}

// 	timeRecord.ID = id
// 	timeRecord.CreatedAt = time.Now()

// 	err := timeRecord.isValid()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &timeRecord, nil
// }

// func (t *TimeRecord) isValid() error {

// 	// TODO: change 5 for company tolerance
// 	if t.Time.After(time.Now().Add(time.Minute * 5)) {
// 		return errors.New("the registration time must not be longer than the current time")
// 	}

// 	_, err := govalidator.ValidateStruct(t)
// 	return err
// }

// func (t *TimeRecord) Approve(approver *Employee) error {

// 	if *t.EmployeeID == approver.ID {
// 		return errors.New("the employee who recorded the time cannot be the same person who approves")
// 	}

// 	if t.Status == TIME_RECORD_APPROVED {
// 		return errors.New("the time record has already been approved")
// 	}

// 	if t.Status == TIME_RECORD_REFUSED {
// 		return errors.New("the refused time record cannot be approved")
// 	}

// 	t.Status = TIME_RECORD_APPROVED
// 	t.UpdatedAt = time.Now()
// 	err := t.isValid()
// 	return err
// }

// func (t *TimeRecord) Refuse(refuser *Employee, refusedReason string) error {

// 	if *t.EmployeeID == refuser.ID {
// 		return errors.New("the employee who recorded the time cannot be the same person who refuses")
// 	}

// 	if t.Status == TIME_RECORD_APPROVED {
// 		return errors.New("the approved time record cannot be refused")
// 	}

// 	if t.Status == TIME_RECORD_REFUSED {
// 		return errors.New("the time record has already been refused")
// 	}

// 	if refusedReason == "" {
// 		return errors.New("the refused reason must not be empty")
// 	}

// 	t.Status = TIME_RECORD_REFUSED
// 	t.UpdatedAt = time.Now()
// 	err := t.isValid()
// 	return err
// }
