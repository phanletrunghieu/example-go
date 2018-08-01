package lendbook

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/phanletrunghieu/example-go/domain"
)

// pgService implmenter for LendBook serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for LendBook service
func (s *pgService) Create(_ context.Context, p *domain.LendBook) error {
	return s.db.Create(p).Error
}

// Update implement Update for LendBook service
func (s *pgService) Update(_ context.Context, p *domain.LendBook) (*domain.LendBook, error) {
	old := domain.LendBook{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	if !p.User_ID.IsZero() {
		old.User_ID = p.User_ID
	}

	if !p.Book_ID.IsZero() {
		old.Book_ID = p.Book_ID
	}

	if !p.From.IsZero() {
		old.From = p.From
	}
	if !p.To.IsZero() {
		old.To = p.To
	}

	return &old, s.db.Save(&old).Error
}

// Find implement Find for LendBook service
func (s *pgService) Find(_ context.Context, p *domain.LendBook) (*domain.LendBook, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for LendBook service
func (s *pgService) FindAll(_ context.Context) ([]domain.LendBook, error) {
	res := []domain.LendBook{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for LendBook service
func (s *pgService) Delete(_ context.Context, p *domain.LendBook) error {
	old := domain.LendBook{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}
