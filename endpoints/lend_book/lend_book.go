package lendbook

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/phanletrunghieu/example-go/domain"
	"github.com/phanletrunghieu/example-go/service"
)

// CreateData data for CreateLend
type CreateData struct {
	Book_ID domain.UUID `json:"book_id"`
	User_ID domain.UUID `json:"user_id"`
	From    time.Time   `json:"from"`
	To      time.Time   `json:"to"`
}

// CreateRequest request struct for CreateLend
type CreateRequest struct {
	LendBook CreateData `json:"lend_book"`
}

// CreateResponse response struct for CreateLend
type CreateResponse struct {
	LendBook domain.LendBook `json:"lend_book"`
}

// StatusCode customstatus code for success create Lend
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a LendBook
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(CreateRequest)
			lendBook = &domain.LendBook{
				Book_ID: req.LendBook.Book_ID,
				User_ID: req.LendBook.User_ID,
				From:    req.LendBook.From,
				To:      req.LendBook.To,
			}
		)

		err := s.LendBookService.Create(ctx, lendBook)
		if err != nil {
			return nil, err
		}

		return CreateResponse{LendBook: *lendBook}, nil
	}
}

// FindRequest request struct for Find a User
type FindRequest struct {
	LendBookID domain.UUID
}

// FindResponse response struct for Find a User
type FindResponse struct {
	LendBook *domain.LendBook `json:"lend_book"`
}

// MakeFindEndPoint make endpoint for find User
func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var lendFind domain.LendBook
		req := request.(FindRequest)
		lendFind.ID = req.LendBookID

		lendBook, err := s.LendBookService.Find(ctx, &lendFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{LendBook: lendBook}, nil
	}
}

// FindAllRequest request struct for FindAll User
type FindAllRequest struct{}

// FindAllResponse request struct for find all User
type FindAllResponse struct {
	LendBooks []domain.LendBook `json:"lend_books"`
}

// MakeFindAllEndpoint make endpoint for find all User
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		lendBooks, err := s.LendBookService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{LendBooks: lendBooks}, nil
	}
}

// UpdateData data for Create
type UpdateData struct {
	ID      domain.UUID `json:"-"`
	Book_ID domain.UUID `json:"book_id"`
	User_ID domain.UUID `json:"user_id"`
	From    time.Time   `json:"from"`
	To      time.Time   `json:"to"`
}

// UpdateRequest request struct for update
type UpdateRequest struct {
	LendBook UpdateData `json:"lend_book"`
}

// UpdateResponse response struct for Create
type UpdateResponse struct {
	LendBook domain.LendBook `json:"lend_book"`
}

// MakeUpdateEndpoint make endpoint for update a User
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(UpdateRequest)
			lendBook = domain.LendBook{
				Model:   domain.Model{ID: req.LendBook.ID},
				Book_ID: req.LendBook.Book_ID,
				User_ID: req.LendBook.User_ID,
				From:    req.LendBook.From,
				To:      req.LendBook.To,
			}
		)

		res, err := s.LendBookService.Update(ctx, &lendBook)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{LendBook: *res}, nil
	}
}

// DeleteRequest request struct for delete a User
type DeleteRequest struct {
	LendBookID domain.UUID
}

// DeleteResponse response struct for Find a User
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for update a User
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			lendBookFind = domain.LendBook{}
			req          = request.(DeleteRequest)
		)
		lendBookFind.ID = req.LendBookID

		err := s.LendBookService.Delete(ctx, &lendBookFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
