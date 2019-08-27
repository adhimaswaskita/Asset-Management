package repositories_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nrepo "github.com/adhimaswaskita/AssetManagement/repositories"
	"github.com/jinzhu/gorm"
)

func TestUpdateManufacture(t *testing.T) {
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

	t.Run("Update Manufacture OK", func(t *testing.T) {
		var ID uint = 1
		manufacture := &nmodel.Manufacture{
			Name:    "New manufacture 1",
			Comment: "Comment manufacture 1",
		}

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		manufactures, err := repository.UpdateManufacture(ID, manufacture)
		if err != nil {
			t.Errorf("This should not be error but have %v", err)
		}
		if manufactures.Name != "New manufacture 1" {
			t.Errorf("Manufacture name should be equal to 'New manufacture 1' but have %v", manufactures.Name)
		}
	})

	t.Run("Update Manufacture NOK", func(t *testing.T) {
		var ID uint = 1
		manufacture := &nmodel.Manufacture{
			Name:    "New manufacture 1",
			Comment: "Comment manufacture 1",
		}

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnError(errors.New("Update manufacture failed"))
		mock.ExpectCommit()

		_, err := repository.UpdateManufacture(ID, manufacture)
		if err == nil {
			t.Errorf("This should not error but no error occured")
		}
	})
}
