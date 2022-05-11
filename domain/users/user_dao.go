package users

import (
	"fmt"

	dateUtils "github.com/ZerepL/bookstore_users-api/utils/date_utils"
	internalErrors "github.com/ZerepL/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *internalErrors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return internalErrors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *internalErrors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return internalErrors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return internalErrors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DateCreated = dateUtils.GetNowString()

	usersDB[user.Id] = user
	return nil
}
