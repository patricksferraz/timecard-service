package model_test

import (
	"math"
	"testing"

	"dev.azure.com/c4ut/TimeClock/_git/timecard-service/domain/model"
	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestModel_TimeRecordStatus(t *testing.T) {

	status := model.TIME_RECORD_PENDING
	require.Equal(t, status.String(), model.TIME_RECORD_PENDING.String())
	status = model.TIME_RECORD_APPROVED
	require.Equal(t, status.String(), model.TIME_RECORD_APPROVED.String())
	status = model.TIME_RECORD_REFUSED
	require.Equal(t, status.String(), model.TIME_RECORD_REFUSED.String())

	otherStatus := model.TimeRecordStatus(faker.RandomInt(int(model.TIME_RECORD_REFUSED)+1, math.MaxInt64))
	require.Equal(t, otherStatus.String(), "")
}
