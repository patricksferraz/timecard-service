package service

import (
	"context"
	"time"

	"dev.azure.com/c4ut/TimeClock/_git/timecard-service/domain/model"
	"dev.azure.com/c4ut/TimeClock/_git/timecard-service/domain/repository"
)

type TimecardService struct {
	AuthRepository     repository.AuthRepositoryInterface
	TimecardRepository repository.TimecardRepositoryInterface
}

func NewTimecardService(authRepository repository.AuthRepositoryInterface, timecardRepository repository.TimecardRepositoryInterface) *TimecardService {
	return &TimecardService{
		AuthRepository:     authRepository,
		TimecardRepository: timecardRepository,
	}
}

func (t *TimecardService) RegisterTimecard(ctx context.Context, startDate, endDate time.Time, companyID string, employeeID string) (*model.Timecard, error) {
	company, err := t.TimecardRepository.FindCompany(ctx, companyID)
	if err != nil {
		return nil, err
	}

	employee, err := t.TimecardRepository.FindEmployee(ctx, employeeID)
	if err != nil {
		return nil, err
	}

	timecard, err := model.NewTimecard(startDate, endDate, company, employee)
	if err != nil {
		return nil, err
	}

	err = t.TimecardRepository.RegisterTimecard(ctx, timecard)
	if err != nil {
		return nil, err
	}

	return timecard, nil
}

func (t *TimecardService) FindTimecard(ctx context.Context, timecardID string) (*model.Timecard, error) {
	timecard, err := t.TimecardRepository.FindTimecard(ctx, timecardID)
	if err != nil {
		return nil, err
	}

	return timecard, nil
}

func (t *TimecardService) AddTimeRecord(ctx context.Context, timeRecordID string, timecardID string) (*model.TimeRecord, error) {
	timecard, err := t.TimecardRepository.FindTimecard(ctx, timecardID)
	if err != nil {
		return nil, err
	}

	// external
	eTimeRecord, err := t.TimecardRepository.FindTimeRecord(ctx, timeRecordID)
	if err != nil {
		return nil, err
	}

	timeRecord, err := model.NewTimeRecord(eTimeRecord.ID, eTimeRecord.Time, eTimeRecord.Status, eTimeRecord.EmployeeID, timecard)
	if err != nil {
		return nil, err
	}

	err = timecard.AddTimeRecord(timeRecord)
	if err != nil {
		return nil, err
	}

	err = t.TimecardRepository.AddTimeRecord(ctx, timeRecord)
	if err != nil {
		return nil, err
	}

	return timeRecord, nil
}

func (t *TimecardService) TimecardPreview(ctx context.Context, timecardID string) (*model.Timecard, error) {
	timecard, err := t.TimecardRepository.FindTimecard(ctx, timecardID)
	if err != nil {
		return nil, err
	}

	timecard.Calculate()
	return timecard, nil
}

func (t *TimecardService) WaitApprovalTimecard(ctx context.Context, timecardID string, auditorID string) (*model.Timecard, error) {
	auditorEmployee, err := t.TimecardRepository.FindEmployee(ctx, auditorID)
	if err != nil {
		return nil, err
	}

	timecard, err := t.TimecardRepository.FindTimecard(ctx, timecardID)
	if err != nil {
		return nil, err
	}

	err = timecard.WaitApproval(auditorEmployee)
	if err != nil {
		return nil, err
	}

	err = t.TimecardRepository.SaveTimecard(ctx, timecard)
	if err != nil {
		return nil, err
	}

	return timecard, nil
}

func (t *TimecardService) ApproveTimecard(ctx context.Context, timecardID string, approverID string) (*model.Timecard, error) {
	approverEmployee, err := t.TimecardRepository.FindEmployee(ctx, approverID)
	if err != nil {
		return nil, err
	}

	timecard, err := t.TimecardRepository.FindTimecard(ctx, timecardID)
	if err != nil {
		return nil, err
	}

	err = timecard.Approve(approverEmployee)
	if err != nil {
		return nil, err
	}

	err = t.TimecardRepository.SaveTimecard(ctx, timecard)
	if err != nil {
		return nil, err
	}

	return timecard, nil
}

func (t *TimecardService) FindTimecardsByCompanyID(ctx context.Context, companyID string) ([]*model.Timecard, error) {
	timecards, err := t.TimecardRepository.FindByCompanyID(ctx, companyID)
	if err != nil {
		return nil, err
	}
	return timecards, nil
}

func (t *TimecardService) FindAllTimecardByEmployeeID(ctx context.Context, employeeID string) ([]*model.Timecard, error) {
	timecards, err := t.TimecardRepository.FindByEmployeeID(ctx, employeeID)
	if err != nil {
		return nil, err
	}
	return timecards, nil
}
