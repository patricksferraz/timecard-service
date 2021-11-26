package service

import (
	"context"
	"time"

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

	companyEmployee, err := entity.NewCompanyEmployee(company.ID, employee.ID)
	if err != nil {
		return err
	}

	err = s.Repository.AddEmployeeToCompany(ctx, companyEmployee)
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

func (s *Service) RegisterTimeRecord(ctx context.Context, id string, _time time.Time, tzOffset int, status int, employeeID, companyID string) error {
	employee, err := s.Repository.FindEmployee(ctx, employeeID)
	if err != nil {
		return err
	}

	company, err := s.Repository.FindCompany(ctx, companyID)
	if err != nil {
		return err
	}

	timeRecord, err := entity.NewTimeRecord(id, _time, tzOffset, status, employee, company)
	if err != nil {
		return err
	}

	err = s.Repository.RegisterTimeRecord(ctx, timeRecord)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ApproveTimeRecord(ctx context.Context, id string) error {
	timeRecord, err := s.Repository.FindTimeRecord(ctx, id)
	if err != nil {
		return err
	}

	err = timeRecord.Approve()
	if err != nil {
		return err
	}

	err = s.Repository.SaveTimeRecord(ctx, timeRecord)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RefuseTimeRecord(ctx context.Context, id string) error {
	timeRecord, err := s.Repository.FindTimeRecord(ctx, id)
	if err != nil {
		return err
	}

	err = timeRecord.Refuse()
	if err != nil {
		return err
	}

	err = s.Repository.SaveTimeRecord(ctx, timeRecord)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateWorkScale(ctx context.Context, id, companyID string) error {
	company, err := s.Repository.FindCompany(ctx, companyID)
	if err != nil {
		return err
	}

	workScale, err := entity.NewWorkScale(id, company)
	if err != nil {
		return err
	}

	err = s.Repository.CreateWorkScale(ctx, workScale)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) FindWorkScale(ctx context.Context, workScaleID string) (*entity.WorkScale, error) {
	workScale, err := s.Repository.FindWorkScale(ctx, workScaleID)
	if err != nil {
		return nil, err
	}
	return workScale, nil
}

func (s *Service) AddClockToWorkScale(ctx context.Context, id string, clockType int, clock, timezone, workScaleID string) error {
	workScale, err := s.Repository.FindWorkScale(ctx, workScaleID)
	if err != nil {
		return err
	}

	c, err := entity.NewClock(id, clock, clockType, timezone, workScale)
	if err != nil {
		return err
	}

	err = s.Repository.CreateClock(ctx, c)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) FindClock(ctx context.Context, workScaleID, clockID string) (*entity.Clock, error) {
	clock, err := s.Repository.FindClock(ctx, workScaleID, clockID)
	if err != nil {
		return nil, err
	}
	return clock, nil
}

func (s *Service) DeleteClock(ctx context.Context, workScaleID, clockID string) error {
	err := s.Repository.DeleteClock(ctx, workScaleID, clockID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateClock(ctx context.Context, clockType int, clock, timezone, workScaleID, clockID string) error {
	c, err := s.Repository.FindClock(ctx, workScaleID, clockID)
	if err != nil {
		return err
	}

	if err = c.SetType(clockType); err != nil {
		return err
	}
	if err = c.SetClock(clock); err != nil {
		return err
	}
	if err = c.SetTimezone(timezone); err != nil {
		return err
	}
	if err = s.Repository.SaveClock(ctx, c); err != nil {
		return err
	}

	return nil
}

func (s *Service) AddWorkScaleToEmployee(ctx context.Context, companyID, employeeID, workScaleID string) error {
	companyEmployee, err := s.Repository.FindCompanyEmployee(ctx, companyID, employeeID)
	if err != nil {
		return err
	}

	workScale, err := s.Repository.FindWorkScale(ctx, workScaleID)
	if err != nil {
		return err
	}

	err = companyEmployee.SetScale(workScale)
	if err != nil {
		return err
	}

	err = s.Repository.SaveCompanyEmployee(ctx, companyEmployee)
	if err != nil {
		return err
	}

	return nil
}
