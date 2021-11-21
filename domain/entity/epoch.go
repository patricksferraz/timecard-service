package entity

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Epoch struct {
	Base          `bson:",inline" valid:"-"`
	StartRecordID *string     `json:"start_record,omitempty" gorm:"column:start_record;not null;unique_index:idx_employee_company_start" valid:"uuid"`
	StartRecord   *TimeRecord `json:"-" valid:"-"`
	EndRecordID   *string     `json:"end_record,omitempty" gorm:"column:end_record;unique_index:idx_employee_company_end" valid:"uuid,optional"`
	EndRecord     *TimeRecord `json:"-" valid:"-"`
	Status        EpochStatus `json:"status" gorm:"column:status;not null" valid:"epochStatus"`
	EmployeeID    *string     `json:"employee_id,omitempty" gorm:"column:employee_id;type:uuid;not null;unique_index:idx_employee_company_start,idx_employee_company_end" bson:"employee_id" valid:"uuid"`
	Employee      *Employee   `json:"-" valid:"-"`
	CompanyID     *string     `json:"company_id,omitempty" gorm:"column:company_id;type:uuid;not null;unique_index:idx_employee_company_start,idx_employee_company_end" bson:"company_id" valid:"uuid"`
	Company       *Company    `json:"-" valid:"-"`
}

func NewEpoch(startRecord *TimeRecord, employee *Employee, company *Company) (*Epoch, error) {
	epoch := Epoch{
		StartRecordID: &startRecord.ID,
		StartRecord:   startRecord,
		Status:        EPOCH_PENDING,
		EmployeeID:    &employee.ID,
		Employee:      employee,
		CompanyID:     &company.ID,
		Company:       company,
	}

	epoch.ID = uuid.NewV4().String()
	epoch.CreatedAt = time.Now()

	err := epoch.isValid()
	if err != nil {
		return nil, err
	}

	return &epoch, nil
}

func (t *Epoch) isValid() error {
	_, err := govalidator.ValidateStruct(t)
	return err
}

func (e *Epoch) Complete() error {

	if e.Status == EPOCH_COMPLETED {
		return errors.New("the epoch has already been completed")
	}

	if e.Status == EPOCH_FAILED {
		return errors.New("the failed epoch cannot be completed")
	}

	e.Status = EPOCH_COMPLETED
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Epoch) Fail() error {

	if e.Status == EPOCH_COMPLETED {
		return errors.New("the completed epoch cannot be failed")
	}

	if e.Status == EPOCH_FAILED {
		return errors.New("the epoch has already been failed")
	}

	e.Status = EPOCH_FAILED
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}
