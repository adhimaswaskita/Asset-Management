package repositories_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nrepo "github.com/adhimaswaskita/AssetManagement/repositories"
	"github.com/jinzhu/gorm"
)

func TestCreateManufacture(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open("sqlite3", db)
	if err != nil {
		t.Errorf("Error open mocked gorm")
	}

	repository := &nrepo.Repository{
		DB: gormDB,
	}

	t.Run("Testing Create Manufacture OK", func(t *testing.T) {
		manufacture := &nmodel.Manufacture{
			Name:    "Manufacture 1",
			Comment: "Comment",
		}

		mock.ExpectBegin()
		mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		manufactures, err := repository.CreateManufacture(manufacture)
		if err != nil {
			t.Errorf("This should not be error")
		}
		if manufactures == nil {
			t.Errorf("This should not be 0")
		}
	})

	t.Run("Create Manufacture NOK", func(t *testing.T) {
		manufacture := &nmodel.Manufacture{
			Name:    "Manufacture 1",
			Comment: "Comment",
		}

		mock.ExpectBegin()
		mock.ExpectExec("INSERT").WillReturnError(errors.New("Insert manufacture failed"))
		mock.ExpectCommit()

		_, err := repository.CreateManufacture(manufacture)
		if err == nil {
			t.Errorf("This should be an error")
		}

	})
}
