package entity

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.TagMap["eventStatus"] = govalidator.Validator(func(str string) bool {
		res := str == EVENT_PENDING.String()
		res = res || str == EVENT_COMPLETED.String()
		res = res || str == EVENT_FAILED.String()
		return res
	})
}

type EventStatus int

const (
	EVENT_PENDING EventStatus = iota + 1
	EVENT_COMPLETED
	EVENT_FAILED
)

func (s EventStatus) String() string {
	switch s {
	case EVENT_PENDING:
		return "PENDING"
	case EVENT_COMPLETED:
		return "COMPLETED"
	case EVENT_FAILED:
		return "FAILED"
	}
	return ""
}
