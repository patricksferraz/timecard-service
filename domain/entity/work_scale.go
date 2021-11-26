package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type WorkScale struct {
	Base      `json:",inline" valid:"required"`
	Clocks    []*Clock `json:"clocks,omitempty" gorm:"ForeignKey:WorkScaleID" valid:"-"`
	CompanyID *string  `json:"company_id,omitempty" gorm:"column:company_id;type:uuid;not null" valid:"uuid"`
	Company   *Company `json:"-" valid:"-"`
}

func NewWorkScale(id string, company *Company, clocks ...*Clock) (*WorkScale, error) {
	entity := &WorkScale{
		Clocks:    clocks,
		CompanyID: &company.ID,
		Company:   company,
	}
	entity.ID = id
	entity.CreatedAt = time.Now()

	err := entity.isValid()
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *WorkScale) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
