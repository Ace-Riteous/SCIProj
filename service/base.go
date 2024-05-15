package service

var baseService *BaseService

type BaseService struct {
}

func NewBaseService() *BaseService {
	if baseService == nil {
		baseService = &BaseService{}
	}
	return baseService
}
