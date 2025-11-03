package errs

import (
	"errors"
	"net/http"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrSQLQuery           = errors.New("sql query problem")
	ErrScanningRows       = errors.New("error while scanning rows")
	ErrUsernameTaken      = errors.New("this username is already taken")
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrTableMissing       = errors.New("table not found")
	ErrHashingPass        = errors.New("error while hashing password")
	ErrInvalidJSON        = errors.New("invalid JSON body")
	ErrCreatingUser       = errors.New("error while creating user")
	ErrMissingTokens      = errors.New("missing tokens")
	ErrInvalidToken       = errors.New("invalid or expired token")
	ErrTokenGeneration    = errors.New("failed to generate tokens")
	ErrUnauthorized       = errors.New("unauthorized access")
)

var ErrorCodeToHTTPStatus = map[error]int{
	ErrUserNotFound:       http.StatusNotFound,
	ErrSQLQuery:           http.StatusInternalServerError,
	ErrScanningRows:       http.StatusInternalServerError,
	ErrUsernameTaken:      http.StatusConflict,
	ErrInvalidCredentials: http.StatusBadRequest,
	ErrTableMissing:       http.StatusInternalServerError,
	ErrHashingPass:        http.StatusInternalServerError,
	ErrInvalidJSON:        http.StatusBadRequest,
	ErrCreatingUser:       http.StatusInternalServerError,
	ErrMissingTokens:      http.StatusUnauthorized,
	ErrInvalidToken:       http.StatusUnauthorized,
	ErrTokenGeneration:    http.StatusInternalServerError,
	ErrUnauthorized:       http.StatusUnauthorized,
}
