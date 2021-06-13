package model_test

import (
	"math"
	"testing"

	"dev.azure.com/c4ut/TimeClock/_git/timecard-service/domain/model"
	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestModel_TimecardStatus(t *testing.T) {

	status := model.TIMECARD_OPEN
	require.Equal(t, status.String(), model.TIMECARD_OPEN.String())
	status = model.TIMECARD_AWAITING_APPROVAL
	require.Equal(t, status.String(), model.TIMECARD_AWAITING_APPROVAL.String())
	status = model.TIMECARD_APPROVED
	require.Equal(t, status.String(), model.TIMECARD_APPROVED.String())

	otherStatus := model.TimecardStatus(faker.RandomInt(int(model.TIMECARD_APPROVED)+1, math.MaxInt64))
	require.Equal(t, otherStatus.String(), "")
}
