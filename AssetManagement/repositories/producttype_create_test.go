package repositories_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nrepo "github.com/adhimaswaskita/AssetManagement/repositories"
	"github.com/jinzhu/gorm"
)

func TestCreateProductType(t *testing.T) {
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

	t.Run("Testing Create Product Type OK", func(t *testing.T) {
		productType := &nmodel.ProductType{
			Name:        "product type 1",
			Type:        "asset",
			Category:    "Non-IT",
			Description: "desciption",
		}

		mock.ExpectBegin()
		mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		productTypes, err := repository.CreateProductType(productType)
		if err != nil {
			t.Errorf("This should not be error")
		}
		if productTypes == nil {
			t.Errorf("This should not be 0")
		}
	})
}
