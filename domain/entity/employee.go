package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Employee struct {
	Base `json:",inline" valid:"required"`
	// TimeRecords []*TimeRecord `json:"-" gorm:"ForeignKey:EmployeeID" valid:"-"`
	Companies []*Company `json:"-" gorm:"many2many:companies_employees" valid:"-"`
}

func NewEmployee(id string) (*Employee, error) {
	entity := &Employee{}
	entity.ID = id
	entity.CreatedAt = time.Now()

	err := entity.isValid()
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *Employee) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *Employee) AddCompany(company *Company) error {
	e.Companies = append(e.Companies, company)
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}
