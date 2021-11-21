package entity

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.TagMap["epochStatus"] = govalidator.Validator(func(str string) bool {
		res := str == EPOCH_PENDING.String()
		res = res || str == EPOCH_FAILED.String()
		res = res || str == EPOCH_COMPLETED.String()
		return res
	})
}

type EpochStatus int

const (
	EPOCH_PENDING EpochStatus = iota + 1
	EPOCH_FAILED
	EPOCH_COMPLETED
)

func (s EpochStatus) String() string {
	switch s {
	case EPOCH_PENDING:
		return "PENDING"
	case EPOCH_FAILED:
		return "FAILED"
	case EPOCH_COMPLETED:
		return "COMPLETED"
	}
	return ""
}
