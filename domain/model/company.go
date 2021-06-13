package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Company struct {
	Base      `valid:"required"`
	Employees []*Employee `gorm:"many2many:company_employees" valid:"-"`
	Timecards []*Timecard `gorm:"column:timecards;foreignKey:CompanyID" valid:"-"`
}

func NewCompany(id string) (*Company, error) {

	company := Company{}

	company.ID = id
	company.CreatedAt = time.Now()

	err := company.isValid()
	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (c *Company) isValid() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}
