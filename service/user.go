package service

import "web-study/repository"

type User struct {
	userRepository repository.UserRepository
}

func newUserService(userRepository *repository.UserRepository) *User {
	return &User{}
}
