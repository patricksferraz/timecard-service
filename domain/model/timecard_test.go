package model_test

import (
	"testing"
	"time"

	"dev.azure.com/c4ut/TimeClock/_git/timecard-service/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewTimecard(t *testing.T) {

	companyID := uuid.NewV4().String()
	employeeID := uuid.NewV4().String()
	startDate := time.Now()
	endDate := time.Now()

	var roles []string
	company, _ := model.NewCompany(companyID)
	employee, _ := model.NewEmployee(employeeID, roles)
	timecard, err := model.NewTimecard(startDate, endDate, company, employee)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(timecard.ID))
	require.Equal(t, timecard.Status, model.TIMECARD_OPEN)
	require.Equal(t, timecard.StartDate, startDate)
	require.Equal(t, timecard.EndDate, endDate)
	require.Equal(t, timecard.CompanyID, companyID)
	require.Equal(t, timecard.EmployeeID, employeeID)

	_, err = model.NewTimecard(time.Time{}, endDate, company, employee)
	require.NotNil(t, err)
	_, err = model.NewTimecard(startDate, time.Time{}, company, employee)
	require.NotNil(t, err)
	_, err = model.NewTimecard(startDate, endDate, &model.Company{}, employee)
	require.NotNil(t, err)
	_, err = model.NewTimecard(startDate, endDate, company, &model.Employee{})
	require.NotNil(t, err)
}

func TestModel_ChangeStatusOfATimecard(t *testing.T) {

	companyID := uuid.NewV4().String()
	employeeID := uuid.NewV4().String()
	auditorID := uuid.NewV4().String()
	approverID := uuid.NewV4().String()
	startDate := time.Now()
	endDate := time.Now()

	var roles []string
	company, _ := model.NewCompany(companyID)
	employee, _ := model.NewEmployee(employeeID, roles)
	auditor, _ := model.NewEmployee(auditorID, roles)
	approver, _ := model.NewEmployee(approverID, roles)

	timecard, _ := model.NewTimecard(startDate, endDate, company, employee)

	err := timecard.WaitApproval(auditor)
	require.Nil(t, err)
	require.Equal(t, timecard.Status, model.TIMECARD_AWAITING_APPROVAL)
	require.True(t, timecard.CreatedAt.Before(timecard.UpdatedAt))

	err = timecard.Approve(approver)
	require.Nil(t, err)
	require.Equal(t, timecard.Status, model.TIMECARD_APPROVED)
	require.True(t, timecard.CreatedAt.Before(timecard.UpdatedAt))
}

func TestModel_AddTimeRecordInATimecard(t *testing.T) {

	companyID := uuid.NewV4().String()
	employeeID := uuid.NewV4().String()
	startDate := time.Now()
	endDate := time.Now()

	var roles []string
	company, _ := model.NewCompany(companyID)
	employee, _ := model.NewEmployee(employeeID, roles)
	timecard, _ := model.NewTimecard(startDate, endDate, company, employee)

	timeRecord, _ := model.NewTimeRecord(
		uuid.NewV4().String(),
		time.Now(),
		model.TIME_RECORD_APPROVED,
		employeeID,
		timecard,
	)
	err := timecard.AddTimeRecord(timeRecord)
	require.Nil(t, err)
	require.Equal(t, timecard.Status, model.TIMECARD_OPEN)
	require.Len(t, len(timecard.TimeRecords), 1)
	require.True(t, timecard.CreatedAt.Before(timecard.UpdatedAt))
}
