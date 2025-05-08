package rest

import (
	"github.com/patricksferraz/timecard-service/domain/service"
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
