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
}
