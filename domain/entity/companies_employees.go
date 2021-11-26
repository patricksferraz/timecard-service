package entity

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type CompaniesEmployee struct {
	// gorm.JoinTableHandlerInterface
	CompanyID   string     `gorm:"column:company_id;type:uuid;not null;unique_index:idx_company_employee_work_scale;primaryKey" valid:"uuid"`
	EmployeeID  string     `gorm:"column:employee_id;type:uuid;not null;unique_index:idx_company_employee_work_scale;primaryKey" valid:"uuid"`
	WorkScaleID *string    `gorm:"column:work_scale_id;type:uuid;unique_index:idx_company_employee_work_scale" valid:"uuid,optional"`
	WorkScale   *WorkScale `json:"-" valid:"-"`
}

func NewCompanyEmployee(companyID, employeeID string) (*CompaniesEmployee, error) {
	entity := &CompaniesEmployee{
		CompanyID:  companyID,
		EmployeeID: employeeID,
	}

	err := entity.isValid()
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *CompaniesEmployee) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *CompaniesEmployee) SetScale(workScale *WorkScale) error {
	e.WorkScaleID = &workScale.ID
	e.WorkScale = workScale
	err := e.isValid()
	return err
}
