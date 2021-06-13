package model

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.TagMap["timecardStatus"] = govalidator.Validator(func(str string) bool {
		res := str == TIMECARD_OPEN.String()
		res = res || str == TIMECARD_AWAITING_APPROVAL.String()
		res = res || str == TIMECARD_APPROVED.String()
		return res
	})
}

type TimecardStatus int

const (
	TIMECARD_OPEN TimecardStatus = iota + 1
	TIMECARD_AWAITING_APPROVAL
	TIMECARD_APPROVED
)

func (o TimecardStatus) String() string {
	switch o {
	case TIMECARD_OPEN:
		return "OPEN"
	case TIMECARD_AWAITING_APPROVAL:
		return "AWAITING APPROVAL"
	case TIMECARD_APPROVED:
		return "APPROVED"
	}
	return ""
}
