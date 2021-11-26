package entity

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Epoch struct {
	Base           `bson:",inline" valid:"-"`
	InputRecordID  *string       `json:"input_record,omitempty" gorm:"column:input_record;not null;unique_index:idx_employee_company_input" valid:"uuid"`
	InputRecord    *TimeRecord   `json:"-" valid:"-"`
	OutputRecordID *string       `json:"output_record,omitempty" gorm:"column:output_record;unique_index:idx_employee_company_output" valid:"uuid,optional"`
	OutputRecord   *TimeRecord   `json:"-" valid:"-"`
	WorkedHours    time.Duration `json:"worked_hours,omitempty" gorm:"column:worked_hours" valid:"-"`
	Status         EpochStatus   `json:"status" gorm:"column:status;not null" valid:"epochStatus"`
	EmployeeID     *string       `json:"employee_id,omitempty" gorm:"column:employee_id;type:uuid;not null;unique_index:idx_employee_company_input,idx_employee_company_output" bson:"employee_id" valid:"uuid"`
	Employee       *Employee     `json:"-" valid:"-"`
	CompanyID      *string       `json:"company_id,omitempty" gorm:"column:company_id;type:uuid;not null;unique_index:idx_employee_company_input,idx_employee_company_output" bson:"company_id" valid:"uuid"`
	Company        *Company      `json:"-" valid:"-"`
	Token          *string       `json:"-" gorm:"column:token;type:varchar(25);not null" valid:"-"`
}

func NewEpoch(inputRecord *TimeRecord, employee *Employee, company *Company) (*Epoch, error) {
	token := primitive.NewObjectID().Hex()
	epoch := Epoch{
		InputRecordID: &inputRecord.ID,
		InputRecord:   inputRecord,
		Status:        EPOCH_PENDING,
		EmployeeID:    &employee.ID,
		Employee:      employee,
		CompanyID:     &company.ID,
		Company:       company,
		Token:         &token,
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

func (e *Epoch) Complete(outputRecord *TimeRecord) error {

	if e.Status == EPOCH_PROCESSED {
		return errors.New("the processed epoch cannot be completed")
	}

	if e.Status == EPOCH_FAILED {
		return errors.New("the failed epoch cannot be completed")
	}

	e.OutputRecordID = &outputRecord.ID
	e.OutputRecord = outputRecord
	e.Status = EPOCH_COMPLETED
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Epoch) Process() error {

	if e.Status == EPOCH_FAILED {
		return errors.New("the failed epoch cannot be processed")
	}

	if e.Status == EPOCH_PROCESSED {
		return errors.New("the epoch has already been processed")
	}

	e.Status = EPOCH_PROCESSED
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Epoch) Fail() error {

	if e.Status == EPOCH_PROCESSED {
		return errors.New("the processed epoch cannot be failed")
	}

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
