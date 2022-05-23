package users

import (
	"github.com/ZerepL/bookstore_users-api/datasource/mysql/users_db"
	dateUtils "github.com/ZerepL/bookstore_users-api/utils/date_utils"
	internalErrors "github.com/ZerepL/bookstore_users-api/utils/errors"
	"github.com/ZerepL/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT * FROM users WHERE id=?"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
)

func (user *User) Get() *internalErrors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return internalErrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if saveErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); saveErr != nil {
		return mysql_utils.ParserError(saveErr)
	}

	return nil
}

func (user *User) Save() *internalErrors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return internalErrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = dateUtils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParserError(saveErr)
	}
	
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParserError(err)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *internalErrors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return internalErrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParserError(err)
	}
	return nil
}
