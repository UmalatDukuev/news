package errs

import (
	"errors"
	"net/http"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrSQLQuery      = errors.New("sql query problem")
	ErrScanningRows  = errors.New("error while scanning rows")
	ErrUsernameTaken = errors.New("this username is already taken")
)

var ErrorCodeToHTTPStatus = map[error]int{
	ErrUserNotFound:  http.StatusNotFound,
	ErrSQLQuery:      http.StatusInternalServerError,
	ErrScanningRows:  http.StatusInternalServerError,
	ErrUsernameTaken: http.StatusConflict,
}
