package repositories_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	nrepo "github.com/adhimaswaskita/AssetManagement/repositories"
	"github.com/jinzhu/gorm"
)

func TestDeleteProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		t.Errorf("Error open mocked gorm")
	}

	repository := &nrepo.Repository{
		DB: gormDB,
	}

	t.Run("Delete Product OK", func(t *testing.T) {
		ID := uint(1)

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		err := repository.DeleteProduct(ID)
		if err != nil {
			t.Errorf("This should not be error, but have %v", err)
		}
	})

	t.Run("Delete Product NOK", func(t *testing.T) {
		ID := uint(1)

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnError(errors.New("Delete product failed"))
		mock.ExpectCommit()

		err := repository.DeleteProduct(ID)
		if err == nil {
			t.Errorf("This should not be error, but have %v", err)
		}
	})
}
