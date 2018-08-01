package lendbook

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound         = errNotFound{}
	ErrUnknown          = errUnknown{}
	ErrNameIsRequired   = errNameIsRequired{}
	ErrRecordNotFound   = errRecordNotFound{}
	ErrBookNotFound     = errBookNotFound{}
	ErrUserNotFound     = errUserNotFound{}
	ErrBookNotAvailable = errBookNotAvailable{}
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
	return "user name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

// book not found
type errBookNotFound struct{}

func (errBookNotFound) Error() string {
	return "book not found"
}

func (errBookNotFound) StatusCode() int {
	return http.StatusBadRequest
}

// user not found
type errUserNotFound struct{}

func (errUserNotFound) Error() string {
	return "user not found"
}

func (errUserNotFound) StatusCode() int {
	return http.StatusBadRequest
}

// book not availble
type errBookNotAvailable struct{}

func (errBookNotAvailable) Error() string {
	return "book not availble"
}

func (errBookNotAvailable) StatusCode() int {
	return http.StatusBadRequest
}
