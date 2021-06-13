package repository

import (
	"context"

	"dev.azure.com/c4ut/TimeClock/_git/timecard-service/domain/model"
)

type AuthRepositoryInterface interface {
	Verify(ctx context.Context, accessToken string) (*model.Employee, error)
}
