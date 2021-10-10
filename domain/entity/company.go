package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Company struct {
	Base      `json:",inline" valid:"required"`
	Employees []*Employee `json:"-" gorm:"many2many:companies_employees" valid:"-"`
}

func NewCompany(id string) (*Company, error) {
	company := Company{}
	company.ID = id
	company.CreatedAt = time.Now()

	if err := company.isValid(); err != nil {
		return nil, err
	}

	return &company, nil
}

func (c *Company) isValid() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}
