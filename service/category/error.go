package category

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound       = errNotFound{}
	ErrUnknown        = errUnknown{}
	ErrNameIsRequired = errNameIsRequired{}
	ErrNameLenght     = errNameLength{}
	ErrNameExist      = errNameExist{}
	ErrRecordNotFound = errRecordNotFound{}
)

// not found
type errNotFound struct{}

func (errNotFound) Error() string {
	return "record not found"
}
func (errNotFound) StatusCode() int {
	return http.StatusNotFound
}

//err unknown
type errUnknown struct{}

func (errUnknown) Error() string {
	return "unknown error"
}
func (errUnknown) StatusCode() int {
	return http.StatusBadRequest
}

//record not found
type errRecordNotFound struct{}

func (errRecordNotFound) Error() string {
	return "client record not found"
}
func (errRecordNotFound) StatusCode() int {
	return http.StatusNotFound
}

// name not empty
type errNameIsRequired struct{}

func (errNameIsRequired) Error() string {
	return "user name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

// length must > 5 character
type errNameLength struct{}

func (errNameLength) Error() string {
	return "length of name must be greater than 5 characters"
}

func (errNameLength) StatusCode() int {
	return http.StatusBadRequest
}

// name is exist
type errNameExist struct{}

func (errNameExist) Error() string {
	return "name is exist"
}

func (errNameExist) StatusCode() int {
	return http.StatusBadRequest
}
