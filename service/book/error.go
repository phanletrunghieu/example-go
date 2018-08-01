package book

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound          = errNotFound{}
	ErrUnknown           = errUnknown{}
	ErrNameIsRequired    = errNameIsRequired{}
	ErrNameLenght        = errNameLength{}
	ErrNameExist         = errNameExist{}
	ErrDescriptionLength = errDescriptionLength{}
	ErrRecordNotFound    = errRecordNotFound{}
	ErrCategoryNotFound  = errCategoryNotFound{}
)

type errNotFound struct{}

func (errNotFound) Error() string {
	return "record not found"
}
func (errNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errUnknown struct{}

func (errUnknown) Error() string {
	return "unknown error"
}
func (errUnknown) StatusCode() int {
	return http.StatusBadRequest
}

type errRecordNotFound struct{}

func (errRecordNotFound) Error() string {
	return "client record not found"
}
func (errRecordNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errNameIsRequired struct{}

func (errNameIsRequired) Error() string {
	return "book name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

// name length must > 5 character
type errNameLength struct{}

func (errNameLength) Error() string {
	return "length of name must be greater than 5 characters"
}

func (errNameLength) StatusCode() int {
	return http.StatusBadRequest
}

// description length must > 5 character
type errDescriptionLength struct{}

func (errDescriptionLength) Error() string {
	return "length of description must be greater than 5 characters"
}

func (errDescriptionLength) StatusCode() int {
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

// category must exist
type errCategoryNotFound struct{}

func (errCategoryNotFound) Error() string {
	return "category not found"
}

func (errCategoryNotFound) StatusCode() int {
	return http.StatusBadRequest
}
