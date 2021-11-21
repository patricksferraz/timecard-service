package entity

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.TagMap["timeRecordStatus"] = govalidator.Validator(func(str string) bool {
		res := str == TIME_RECORD_PENDING.String()
		res = res || str == TIME_RECORD_APPROVED.String()
		res = res || str == TIME_RECORD_REFUSED.String()
		return res
	})
}

type TimeRecordStatus int

const (
	TIME_RECORD_PENDING TimeRecordStatus = iota + 1
	TIME_RECORD_APPROVED
	TIME_RECORD_REFUSED
)

func (t TimeRecordStatus) String() string {
	switch t {
	case TIME_RECORD_PENDING:
		return "PENDING"
	case TIME_RECORD_APPROVED:
		return "APPROVED"
	case TIME_RECORD_REFUSED:
		return "REFUSED"
	}
	return ""
}
