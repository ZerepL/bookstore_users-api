package mysql_utils

import (
	internalErrors "github.com/ZerepL/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParserError(err error) *internalErrors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return internalErrors.NewNotFoundError("no record matching given id")
		}

		return internalErrors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return internalErrors.NewBadRequestError("duplicated data")
	}

	return internalErrors.NewInternalServerError("error processing request")
}
