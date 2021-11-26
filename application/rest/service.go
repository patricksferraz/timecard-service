package rest

import (
	"github.com/c-4u/timecard-service/domain/service"
)

type RestService struct {
	Service        *service.Service
	AuthMiddleware *AuthMiddleware
}

func NewRestService(service *service.Service, authMiddleware *AuthMiddleware) *RestService {
	return &RestService{
		Service:        service,
		AuthMiddleware: authMiddleware,
	}
}
