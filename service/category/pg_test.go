package category

import (
	"context"
	"testing"

	testutil "github.com/phanletrunghieu/example-go/config/database/pg/util"
	"github.com/phanletrunghieu/example-go/domain"
)

func Test_pgService_Create(t *testing.T) {
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	type args struct {
		p *domain.Category
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				&domain.Category{
					Name: "Create New Category 1",
				},
			},
		},
		{
			name: "Create an exist category",
			args: args{
				&domain.Category{
					Name: "Create New Category 1",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			if err := s.Create(context.Background(), tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("pgService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pgService_Update(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	category1 := domain.Category{
		Name: "Category 1",
	}
	category2 := domain.Category{
		Name: "Category 2",
	}
	err1 := testDB.Create(&category1).Error
	err2 := testDB.Create(&category2).Error
	if err1 != nil || err2 != nil {
		t.Fatalf("Failed to create category")
	}

	fakeID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Category
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				&domain.Category{
					Model: domain.Model{ID: category1.ID},
					Name:  "Category 3",
				},
			},
		},
		{
			name: "fake category id",
			args: args{
				&domain.Category{
					Model: domain.Model{ID: fakeID},
					Name:  "Category 1",
				},
			},
			wantErr: true,
		},
		{
			name: "duplicate name",
			args: args{
				&domain.Category{
					Model: domain.Model{ID: category1.ID},
					Name:  "Category 2",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			if _, err := s.Update(context.Background(), tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("pgService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pgService_Delete(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	category := domain.Category{}
	err = testDB.Create(&category).Error
	if err != nil {
		t.Fatalf("Failed to create category by error %v", err)
	}

	fakeID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Category
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success delete",
			args: args{
				&domain.Category{
					Model: domain.Model{ID: category.ID},
				},
			},
		},
		{
			name: "failed delete by not exist category id",
			args: args{
				&domain.Category{
					Model: domain.Model{ID: fakeID},
				},
			},
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			err := s.Delete(context.Background(), tt.args.p)
			if err != nil && err != tt.wantErr {
				t.Errorf("pgService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && tt.wantErr != nil {
				t.Errorf("pgService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPGService_Find(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	category := domain.Category{}
	err = testDB.Create(&category).Error
	if err != nil {
		t.Fatalf("Failed to create category by error %v", err)
	}

	fakeID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Category
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Category
		wantErr error
	}{
		{
			name: "success find correct category",
			args: args{
				&domain.Category{
					Model: domain.Model{ID: category.ID},
				},
			},
			want: &domain.Category{
				Model: domain.Model{ID: category.ID},
			},
		},
		{
			name: "failed find category by not exist category id",
			args: args{
				&domain.Category{
					Model: domain.Model{ID: fakeID},
				},
			},
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}

			got, err := s.Find(context.Background(), tt.args.p)
			if err != nil && err != tt.wantErr {
				t.Errorf("pgService.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && tt.wantErr != nil {
				t.Errorf("pgService.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && got.ID.String() != tt.want.ID.String() {
				t.Errorf("pgService.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPGService_FindAll(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	category := domain.Category{}
	err = testDB.Create(&category).Error
	if err != nil {
		t.Fatalf("Failed to create category by error %v", err)
	}

	tests := []struct {
		name    string
		want    *domain.Category
		wantErr bool
	}{
		{
			name: "find all",
			want: &domain.Category{Model: domain.Model{ID: category.ID}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			got, err := s.FindAll(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("pgService.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) <= 0 || got[0].ID.String() != tt.want.ID.String() {
				t.Errorf("pgService.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
