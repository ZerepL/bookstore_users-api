package users

import (
	"fmt"
	"github.com/ZerepL/bookstore_users-api/logger"

	"github.com/ZerepL/bookstore_users-api/datasource/mysql/users_db"
	internalErrors "github.com/ZerepL/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
)

func (user *User) Get() *internalErrors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return internalErrors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if saveErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); saveErr != nil {
		logger.Error("error when trying to get user by id", saveErr)
		//return mysql_utils.ParserError(saveErr) // TODO: Convert this to use log
		return internalErrors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Save() *internalErrors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return internalErrors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		//return mysql_utils.ParserError(saveErr)  // TODO: Convert this to use log
		logger.Error("error when trying to save user", err)
		return internalErrors.NewInternalServerError("database error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		//return mysql_utils.ParserError(err) // TODO: Convert this to use log
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return internalErrors.NewInternalServerError("database error")
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *internalErrors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return internalErrors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		//return mysql_utils.ParserError(err) // TODO: Convert this to use log
		logger.Error("error when trying to update user", err)
		return internalErrors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Delete() *internalErrors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return internalErrors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		//return mysql_utils.ParserError(err) // TODO: Convert this to use log
		logger.Error("error when trying to delete user", err)
		return internalErrors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *internalErrors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, internalErrors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find user by status", err)
		return nil, internalErrors.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, internalErrors.NewInternalServerError("database error")
			//return nil, mysql_utils.ParserError(err)  // TODO: Convert this to use log
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, internalErrors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}

func (user *User) FindByEmailAndPassword() *internalErrors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare find user by email and password statement", err)
		return internalErrors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if saveErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); saveErr != nil {
		logger.Error("error when trying to get user by email and password", saveErr)
		//return mysql_utils.ParserError(saveErr) // TODO: Convert this to use log
		//TODO: Parse error for invalid credential
		return internalErrors.NewInternalServerError("database error")
	}
	return nil
}
