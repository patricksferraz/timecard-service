package model_test

import (
	"testing"
	"time"

	"dev.azure.com/c4ut/TimeClock/_git/timecard-service/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewTimeRecord(t *testing.T) {

	id := uuid.NewV4().String()
	employeeID := uuid.NewV4().String()
	companyID := uuid.NewV4().String()
	startDate := time.Now()
	endDate := time.Now()

	var roles []string
	company, _ := model.NewCompany(companyID)
	employee, _ := model.NewEmployee(employeeID, roles)
	timecard, _ := model.NewTimecard(startDate, endDate, company, employee)

	status := model.TIME_RECORD_PENDING
	_time := time.Now()
	timeRecord, err := model.NewTimeRecord(id, _time, status, employeeID, timecard)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(timeRecord.ID))
	require.Equal(t, timeRecord.Time, _time)
	require.Equal(t, timeRecord.Status, status)
	require.Equal(t, timeRecord.EmployeeID, employeeID)

	_, err = model.NewTimeRecord("", _time, status, employeeID, timecard)
	require.NotNil(t, err)
	_, err = model.NewTimeRecord(id, time.Time{}, status, employeeID, timecard)
	require.NotNil(t, err)
	_, err = model.NewTimeRecord(id, _time, 0, employeeID, timecard)
	require.NotNil(t, err)
	_, err = model.NewTimeRecord(id, _time, status, "", &model.Timecard{})
	require.NotNil(t, err)
}
