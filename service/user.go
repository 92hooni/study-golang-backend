package service

import (
	"web-study/repository"
	"web-study/types"
)

type User struct {
	userRepository repository.UserRepository
}

func newUserService(userRepository *repository.UserRepository) *User {
	return &User{}
}

func (u *User) Create(newUser *types.User) error {
	return u.userRepository.Create(newUser)
}

func (u *User) Update(name string, newAge int64) error {
	return u.userRepository.Update(name, newAge)
}

func (u *User) Delete(userName string) error {
	return u.userRepository.Delete(userName)
}

func (u *User) Get() []*types.User {
	return u.userRepository.Get()
}
