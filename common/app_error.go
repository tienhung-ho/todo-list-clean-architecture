package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type appError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *appError {
	return &appError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(root error, msg, log, key string) *appError {
	return &appError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, log, key string) *appError {
	return &appError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func ErrDB(err error) *appError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *appError {
	return NewErrorResponse(err, "Invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternal(err error) *appError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong with the server", err.Error(), "ErrInternal")
}

func ErrCannotListEntity(entity string, err error) *appError {
	return NewErrorResponse(err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity), entity)
}

func ErrCannotGetEntity(entity string, err error) *appError {
	return NewErrorResponse(err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity), entity)
}

func ErrCannotUpdateEntity(entity string, err error) *appError {
	return NewErrorResponse(err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity), entity)
}

func ErrCannotDeleteEntity(entity string, err error) *appError {
	return NewErrorResponse(err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity), entity)
}

func ErrEntityDeleted(entity string, err error) *appError {
	return NewErrorResponse(err,
		fmt.Sprintf("Record has been deleted %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrRecordHasBeenDeleted%s", entity), entity)
}

func ErrCannotCreateEntity(entity string, err error) *appError {
	return NewErrorResponse(err,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity), entity)
}

func ErrNoPermission(entity string, err error) *appError {
	return NewErrorResponse(err,
		"You have no permission",
		"ErrNoPermission", entity)
}

var RecordNotFound = errors.New("record not found!")

func ErrNotFoundEntity(entity string, err error) *appError {
	return NewErrorResponse(err,
		fmt.Sprintf("Cannot not found %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotNotFound%s", entity), entity)
}

func (e *appError) RootError() error {
	if err, ok := e.RootErr.(*appError); ok {
		return err.RootError()
	}
	return e.RootErr
}

// Error implements the error interface for appError
func (e *appError) Error() string {
	return e.RootError().Error()
}
