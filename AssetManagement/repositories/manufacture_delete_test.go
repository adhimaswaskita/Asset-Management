package repositories_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	nrepo "github.com/adhimaswaskita/AssetManagement/repositories"
	"github.com/jinzhu/gorm"
)

func TestDeleteManufacture(t *testing.T) {
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

	t.Run("Delete Manufacture OK", func(t *testing.T) {
		ID := uint(1)

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		err := repository.DeleteManufacture(ID)
		if err != nil {
			t.Errorf("This should not be error, but have %v", err)
		}
	})
}
