package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Employee struct {
	Base      `json:",inline" valid:"required"`
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

// func (e *Employee) AddCompany(company *Company) error {
// 	e.Companies = append(e.Companies, company)
// 	e.UpdatedAt = time.Now()
// 	err := e.isValid()
// 	return err
// }

// func (e *Employee) AddScale(scale *Scale) error {
// 	e.Scales = append(e.Scales, scale)
// 	e.UpdatedAt = time.Now()
// 	err := e.isValid()
// 	return err
// }

// func (e *Employee) GetCompany(companyID string) (*Company, error) {
// 	for _, company := range e.Companies {
// 		if company.ID == companyID {
// 			return company, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("employee does not belong to the company %s", companyID)
// }
