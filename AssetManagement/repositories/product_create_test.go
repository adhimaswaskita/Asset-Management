package repositories_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nrepo "github.com/adhimaswaskita/AssetManagement/repositories"
	"github.com/jinzhu/gorm"
)

func TestCreateProduct(t *testing.T) {
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

	t.Run("Testing Create Product  OK", func(t *testing.T) {
		product := &nmodel.Product{
			Name:              "Product 1",
			ManufactureID:     1,
			ProductTypeID:     1,
			ProductSupplierID: 1,
			PartNo:            19,
		}

		mock.ExpectBegin()
		mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		products, err := repository.CreateProduct(product)
		if err != nil {
			t.Errorf("This should not be error")
		}
		if products == nil {
			t.Errorf("This should not be 0")
		}
	})
	t.Run("Create Product NOK", func(t *testing.T) {
		product := &nmodel.Product{
			Name:              "Product 1",
			ManufactureID:     1,
			ProductTypeID:     1,
			ProductSupplierID: 1,
			PartNo:            19,
		}

		mock.ExpectBegin()
		mock.ExpectExec("INSERT").WillReturnError(errors.New("Insert product failed!"))
		mock.ExpectCommit()

		_, err := repository.CreateProduct(product)
		if err == nil {
			t.Errorf("Create product should be failed")
		}
	})
}
