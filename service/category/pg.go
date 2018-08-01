package category

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/phanletrunghieu/example-go/domain"
)

// pgService implmenter for User serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for User service
func (s *pgService) Create(_ context.Context, p *domain.Category) error {
	find := domain.Category{}
	if err := s.db.Find(&find, "name = ?", p.Name).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			//if not found in db
			return s.db.Create(p).Error
		}

		return err
	}
	return ErrNameExist
}

// Update implement Update for User service
func (s *pgService) Update(_ context.Context, p *domain.Category) (*domain.Category, error) {
	find := domain.Category{}
	if err := s.db.Find(&find, "name = ?", p.Name).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			//if not found in db
			old := domain.Category{Model: domain.Model{ID: p.ID}}
			if err := s.db.Find(&old).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return nil, ErrNotFound
				}
				return nil, err
			}

			old.Name = p.Name

			return &old, s.db.Save(&old).Error
		}

		return nil, err
	}

	return nil, ErrNameExist
}

// Find implement Find for User service
func (s *pgService) Find(_ context.Context, p *domain.Category) (*domain.Category, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for User service
func (s *pgService) FindAll(_ context.Context) ([]domain.Category, error) {
	res := []domain.Category{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for User service
func (s *pgService) Delete(_ context.Context, p *domain.Category) error {
	oldBook := domain.Book{}
	if err := s.db.Find(&oldBook, "category_id = ?", p.ID).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
	}

	if err := s.db.Delete(oldBook).Error; err != nil {
		return err
	}

	old := domain.Category{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}
