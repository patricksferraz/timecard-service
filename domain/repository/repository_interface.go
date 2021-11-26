package repository

import (
	"context"

	"github.com/c-4u/timecard-service/domain/entity"
)

type RepositoryInterface interface {
	CreateEmployee(ctx context.Context, employee *entity.Employee) error
	FindEmployee(ctx context.Context, id string) (*entity.Employee, error)
	SaveEmployee(ctx context.Context, employee *entity.Employee) error

	CreateCompany(ctx context.Context, company *entity.Company) error
	FindCompany(ctx context.Context, id string) (*entity.Company, error)

	CreateEvent(ctx context.Context, event *entity.Event) error
	FindEvent(ctx context.Context, id string) (*entity.Event, error)
	SaveEvent(ctx context.Context, event *entity.Event) error

	PublishEvent(ctx context.Context, msg, topic, key string) error

	RegisterTimeRecord(ctx context.Context, timeRecord *entity.TimeRecord) error
	SaveTimeRecord(ctx context.Context, timeRecord *entity.TimeRecord) error
	FindTimeRecord(ctx context.Context, id string) (*entity.TimeRecord, error)

	CreateEpoch(ctx context.Context, epoch *entity.Epoch) error
	FindEpoch(ctx context.Context, id string) (*entity.Epoch, error)
	SaveEpoch(ctx context.Context, epoch *entity.Epoch) error
	// SearchEpochs(ctx context.Context, filter *entity.Filter) (*string, []*entity.Epoch, error)

	AddEmployeeToCompany(ctx context.Context, companyEmployee *entity.CompaniesEmployee) error
	FindCompanyEmployee(ctx context.Context, companyID, employeeID string) (*entity.CompaniesEmployee, error)
	SaveCompanyEmployee(ctx context.Context, companyEmployee *entity.CompaniesEmployee) error

	CreateWorkScale(ctx context.Context, workScale *entity.WorkScale) error
	FindWorkScale(ctx context.Context, workScaleID string) (*entity.WorkScale, error)
	SaveWorkScale(ctx context.Context, workScale *entity.WorkScale) error

	CreateClock(ctx context.Context, clock *entity.Clock) error
	FindClock(ctx context.Context, workScaleID, clockID string) (*entity.Clock, error)
	DeleteClock(ctx context.Context, workScaleID, clockID string) error
	SaveClock(ctx context.Context, clock *entity.Clock) error
}
