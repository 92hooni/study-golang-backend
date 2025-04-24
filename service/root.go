package service

import (
	"sync"
	"web-study/repository"
)

// network 에서 받은 요청을 repository 등의 데이터를 가공하거나 분석하여 다시 network 로 응답을 내려주는 서비스 로직을 구현한다.

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	// repository
	repository *repository.Repository

	User *User
}

func NewService(rep *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: rep,
		}
	})

	serviceInstance.User = newUserService(rep.User)

	return serviceInstance
}
