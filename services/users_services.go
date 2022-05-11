package services

import (
	"github.com/ZerepL/bookstore_users-api/domain/users"
	internalErrors "github.com/ZerepL/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *internalErrors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func CreateUser(user users.User) (*users.User, *internalErrors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
