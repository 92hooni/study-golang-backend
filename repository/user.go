package repository

import (
	"web-study/types"
	"web-study/types/errors"
)

type UserRepository struct {
	userMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{},
	}
}

func (u *UserRepository) Create(newUser *types.User) error {
	u.userMap = append(u.userMap, newUser)
	return nil
}

func (u *UserRepository) Update(beforeUser *types.User, afterUser *types.User) error {
	return nil
}

func (u *UserRepository) Delete(userName string) error {
	isExisted := false

	for index, user := range u.userMap {
		if user.Name == userName {
			u.userMap = append(u.userMap[0:index], u.userMap[index+1:]...)
			isExisted = true
			break
		}
	}

	if !isExisted {
		return errors.Errorf(errors.NotFoundUser, nil)
	} else {
		return nil
	}
}

func (u *UserRepository) Get() []*types.User {
	return u.userMap
}
