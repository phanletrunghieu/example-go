package service

import "github.com/phanletrunghieu/example-go/service/user"

// Service define list of all services in projects
type Service struct {
	UserService user.Service
}
