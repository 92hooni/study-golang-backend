package repository

import (
	"sync"
)

// DB를 연결하여 사용하는 부분
// 해당 강의에서는 DB를 사용하지 않고 세션으로 대체 할 예정

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	// db mongo.database 대신 Mock 데이터를 활용
	User *UserRepository
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			User: newUserRepository(),
		}
	})

	return repositoryInstance
}
