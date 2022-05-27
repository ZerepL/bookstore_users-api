package services

import (
	"github.com/ZerepL/bookstore_users-api/domain/users"
	"github.com/ZerepL/bookstore_users-api/utils/crypto_utils"
	dateUtils "github.com/ZerepL/bookstore_users-api/utils/date_utils"
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

	user.Status = users.StatusActive
	user.DateCreated = dateUtils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *internalErrors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err = current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteUser(userId int64) *internalErrors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func Search(status string) (users.Users, *internalErrors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
