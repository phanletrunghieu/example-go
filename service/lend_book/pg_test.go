package lendbook

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

	fakeID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")
	category := domain.Category{
		Name: "Category 1",
	}
	book := domain.Book{
		Name: "Book 1",
	}
	user1 := domain.User{
		Name: "User 1",
	}
	user2 := domain.User{
		Name: "User 1",
	}
	err1 := testDB.Create(&category).Error
	err2 := testDB.Create(&book).Error
	err3 := testDB.Create(&user1).Error
	err4 := testDB.Create(&user2).Error
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		t.Fatalf("Failed to create category")
	}

	type args struct {
		p *domain.LendBook
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				&domain.LendBook{
					User_ID: user1.ID,
					Book_ID: book.ID,
				},
			},
		},
		{
			name: "Book not available",
			args: args{
				&domain.LendBook{
					User_ID: user2.ID,
					Book_ID: book.ID,
				},
			},
			wantErr: true,
		},
		{
			name: "Book not exist",
			args: args{
				&domain.LendBook{
					User_ID: user1.ID,
					Book_ID: fakeID,
				},
			},
			wantErr: true,
		},
		{
			name: "User not exist",
			args: args{
				&domain.LendBook{
					User_ID: fakeID,
					Book_ID: book.ID,
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

	fakeID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")
	category := domain.Category{
		Name: "Category 1",
	}
	book1 := domain.Book{
		Name: "Book 1",
	}
	book2 := domain.Book{
		Name: "Book 2",
	}
	user1 := domain.User{
		Name: "User 1",
	}
	user2 := domain.User{
		Name: "User 2",
	}
	lendbook1 := domain.LendBook{
		User_ID: user1.ID,
		Book_ID: book1.ID,
	}
	lendbook2 := domain.LendBook{
		User_ID: user2.ID,
		Book_ID: book2.ID,
	}
	err1 := testDB.Create(&category).Error
	err2 := testDB.Create(&book1).Error
	err3 := testDB.Create(&book2).Error
	err4 := testDB.Create(&user1).Error
	err5 := testDB.Create(&user2).Error
	err6 := testDB.Create(&lendbook1).Error
	err7 := testDB.Create(&lendbook2).Error
	if err1 != nil ||
		err2 != nil ||
		err3 != nil ||
		err4 != nil ||
		err5 != nil ||
		err6 != nil ||
		err7 != nil {
		t.Fatalf("Failed to create data")
	}

	type args struct {
		p *domain.LendBook
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				&domain.LendBook{
					Model:   domain.Model{ID: lendbook1.ID},
					Book_ID: book2.ID,
				},
			},
		},
		{
			name: "book not exist",
			args: args{
				&domain.LendBook{
					Model:   domain.Model{ID: lendbook1.ID},
					Book_ID: fakeID,
				},
			},
			wantErr: true,
		},
		{
			name: "user not exist",
			args: args{
				&domain.LendBook{
					Model:   domain.Model{ID: lendbook1.ID},
					User_ID: fakeID,
				},
			},
			wantErr: true,
		},
		{
			name: "book not available",
			args: args{
				&domain.LendBook{
					Model:   domain.Model{ID: lendbook1.ID},
					Book_ID: book2.ID,
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

	lendbook := domain.LendBook{}
	err = testDB.Create(&lendbook).Error
	if err != nil {
		t.Fatalf("Failed to create lendbook by error %v", err)
	}

	fakeID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.LendBook
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success delete",
			args: args{
				&domain.LendBook{
					Model: domain.Model{ID: lendbook.ID},
				},
			},
		},
		{
			name: "failed delete by not exist lendbook id",
			args: args{
				&domain.LendBook{
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

	lendbook := domain.LendBook{}
	err = testDB.Create(&lendbook).Error
	if err != nil {
		t.Fatalf("Failed to create lendbook by error %v", err)
	}

	fakeID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.LendBook
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.LendBook
		wantErr error
	}{
		{
			name: "success find correct lendbook",
			args: args{
				&domain.LendBook{
					Model: domain.Model{ID: lendbook.ID},
				},
			},
			want: &domain.LendBook{
				Model: domain.Model{ID: lendbook.ID},
			},
		},
		{
			name: "failed find lendbook by not exist lendbook id",
			args: args{
				&domain.LendBook{
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

	lendbook := domain.LendBook{}
	err = testDB.Create(&lendbook).Error
	if err != nil {
		t.Fatalf("Failed to create lendbook by error %v", err)
	}

	tests := []struct {
		name    string
		want    *domain.LendBook
		wantErr bool
	}{
		{
			name: "find all",
			want: &domain.LendBook{Model: domain.Model{ID: lendbook.ID}},
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
