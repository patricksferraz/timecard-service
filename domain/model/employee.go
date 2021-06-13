package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Employee struct {
	Base              `valid:"required"`
	Roles             []string    `valid:"-"`
	Companies         []*Company  `gorm:"many2many:company_employees" valid:"-"`
	SelfTimecards     []*Timecard `gorm:"column:self_timecards;foreignKey:EmployeeID" valid:"-"`
	AuditedTimecards  []*Timecard `gorm:"column:audited_timecards;foreignKey:AuditedBy" valid:"-"`
	ApprovedTimecards []*Timecard `gorm:"column:approved_timecards;foreignKey:ApprovedBy" valid:"-"`
}

func NewEmployee(id string, roles []string) (*Employee, error) {

	employee := Employee{
		Roles: roles,
	}

	employee.ID = id
	employee.CreatedAt = time.Now()

	err := employee.isValid()
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (e *Employee) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
