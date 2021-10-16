package service

import (
	"context"

	"github.com/c-4u/timecard-service/domain/entity"
	"github.com/c-4u/timecard-service/domain/repository"
)

type Service struct {
	Repository repository.RepositoryInterface
}

func NewService(repository repository.RepositoryInterface) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) CreateCompany(ctx context.Context, id string) error {
	company, err := entity.NewCompany(id)
	if err != nil {
		return err
	}

	err = s.Repository.CreateCompany(ctx, company)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateEmployee(ctx context.Context, id string) error {
	employee, err := entity.NewEmployee(id)
	if err != nil {
		return err
	}

	err = s.Repository.CreateEmployee(ctx, employee)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AddEmployeeToCompany(ctx context.Context, companyID, employeeID string) error {
	company, err := s.Repository.FindCompany(ctx, companyID)
	if err != nil {
		return err
	}

	employee, err := s.Repository.FindEmployee(ctx, employeeID)
	if err != nil {
		return err
	}

	err = employee.AddCompany(company)
	if err != nil {
		return err
	}

	err = s.Repository.SaveEmployee(ctx, employee)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ProcessEvent(ctx context.Context, id, resume string) (*entity.Event, error) {
	event, err := s.Repository.FindEvent(ctx, id)
	if err != nil {
		event, err = entity.NewEvent(id, resume)
		if err != nil {
			return nil, err
		}

		err := s.Repository.CreateEvent(ctx, event)
		if err != nil {
			return nil, err
		}

		return event, nil
	}

	err = event.AddAttempt()
	if err != nil {
		s.Repository.SaveEvent(ctx, event)
		return nil, err
	}

	err = s.Repository.SaveEvent(ctx, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *Service) CompleteEvent(ctx context.Context, id string) error {
	event, err := s.Repository.FindEvent(ctx, id)
	if err != nil {
		return err
	}

	err = event.Complete()
	if err != nil {
		return err
	}

	err = s.Repository.SaveEvent(ctx, event)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) FailEvent(ctx context.Context, id string) error {
	event, err := s.Repository.FindEvent(ctx, id)
	if err != nil {
		return err
	}

	err = event.Fail()
	if err != nil {
		return err
	}

	err = s.Repository.SaveEvent(ctx, event)
	if err != nil {
		return err
	}

	return nil
}
