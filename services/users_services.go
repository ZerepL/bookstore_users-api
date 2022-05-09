package services

import (
	"github.com/ZerepL/bookstore_users-api/domain/users"
	internalErrors "github.com/ZerepL/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *internalErrors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return nil, nil
}
