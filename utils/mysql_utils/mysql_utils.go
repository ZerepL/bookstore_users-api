package mysql_utils

import (
	"errors"
	internalErrors "github.com/ZerepL/bookstore_utils/internal_errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) internalErrors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return internalErrors.NewNotFoundError("no record matching given id")
		}
		return internalErrors.NewInternalServerError("error parsing database response", err)
	}

	switch sqlErr.Number {
	case 1062:
		return internalErrors.NewBadRequestError("invalid data")
	}
	return internalErrors.NewInternalServerError("error processing request", errors.New("database error"))
}
