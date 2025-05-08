package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/timecard-service/utils"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)

	govalidator.TagMap["clockType"] = govalidator.Validator(func(str string) bool {
		res := str == CLOCK_INPUT.String()
		res = res || str == CLOCK_OUTPUT.String()
		return res
	})

	govalidator.TagMap["clock"] = govalidator.Validator(func(str string) bool {
		return utils.IsClock(&str)
	})

	govalidator.TagMap["timezone"] = govalidator.Validator(func(str string) bool {
		_, err := time.LoadLocation(str)
		return err == nil
	})
}

type ClockType int

const (
	CLOCK_INPUT ClockType = iota + 1
	CLOCK_OUTPUT
)

func (t ClockType) String() string {
	switch t {
	case CLOCK_INPUT:
		return "INPUT"
	case CLOCK_OUTPUT:
		return "OUTPUT"
	}
	return ""
}

type Clock struct {
	Base        `json:",inline" valid:"required"`
	Type        ClockType  `json:"type" gorm:"column:type;not null;unique_index:idx_clock_type_tz" valid:"clockType"`
	Clock       string     `json:"clock" gorm:"column:clock;not null;unique_index:idx_clock_type_tz" valid:"clock,required"`
	Timezone    string     `json:"timezone" gorm:"column:timezone;not null;unique_index:idx_clock_type_tz" valid:"timezone,required"`
	WorkScaleID string     `json:"work_scale_id" gorm:"column:work_scale_id;type:uuid;not null" valid:"uuid"`
	WorkScale   *WorkScale `json:"-" valid:"-"`
}

func NewClock(id, clock string, clockType int, timezone string, workScale *WorkScale) (*Clock, error) {
	utils.CleanNonDigits(&clock)
	entity := &Clock{
		Clock:       clock,
		Type:        ClockType(clockType),
		Timezone:    timezone,
		WorkScaleID: workScale.ID,
		WorkScale:   workScale,
	}
	entity.ID = id
	entity.CreatedAt = time.Now()

	err := entity.isValid()
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *Clock) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *Clock) SetType(clockType int) error {
	e.Type = ClockType(clockType)
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Clock) SetClock(clock string) error {
	e.Clock = clock
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Clock) SetTimezone(timezone string) error {
	e.Timezone = timezone
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}
