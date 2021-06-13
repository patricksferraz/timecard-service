package repository

import (
	"context"

	"dev.azure.com/c4ut/TimeClock/_git/timecard-service/domain/model"
)

type TimecardRepositoryInterface interface {
	RegisterTimecard(ctx context.Context, timecard *model.Timecard) error
	FindTimecard(ctx context.Context, timecardID string) (*model.Timecard, error)
	SaveTimecard(ctx context.Context, timecard *model.Timecard) error
	FindByCompanyID(ctx context.Context, companyID string) ([]*model.Timecard, error)
	FindByEmployeeID(ctx context.Context, employeeID string) ([]*model.Timecard, error)
	FindTimeRecord(ctx context.Context, timeRecordID string) (*model.TimeRecord, error)
	AddTimeRecord(ctx context.Context, timeRecord *model.TimeRecord) error
	FindEmployee(ctx context.Context, employeeID string) (*model.Employee, error)
	FindCompany(ctx context.Context, companyID string) (*model.Company, error)
}
