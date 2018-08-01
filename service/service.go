package service

import (
	"github.com/phanletrunghieu/example-go/service/book"
	"github.com/phanletrunghieu/example-go/service/category"
	"github.com/phanletrunghieu/example-go/service/lend_book"
	"github.com/phanletrunghieu/example-go/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
	BookService     book.Service
	LendBookService lendbook.Service
}
