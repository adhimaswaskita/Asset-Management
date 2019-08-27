package repositories_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nrepo "github.com/adhimaswaskita/AssetManagement/repositories"
	"github.com/jinzhu/gorm"
)

func TestCreateProductSupplier(t *testing.T) {
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

	t.Run("Testing Create Product Supplier OK", func(t *testing.T) {
		productSupplier := &nmodel.ProductSupplier{
			Name:        "Product Supplier 1",
			Address:     "Product Supplier address 1",
			City:        "Product Supplier city 1",
			State:       "Product Supplier state 1",
			Zip:         68215,
			ContactName: "Product Supplier contact name 1",
			Phone:       "081331555666",
			Email:       "productSupplier1@gmail.com",
			URL:         "www.productsupplier1.io",
		}

		mock.ExpectBegin()
		mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		productSuppliers, err := repository.CreateProductSupplier(productSupplier)
		if err != nil {
			t.Errorf("This should not be error")
		}
		if productSuppliers == nil {
			t.Errorf("This should not be 0")
		}
	})

	t.Run("Create Product Supplier NOK", func(t *testing.T) {
		productSupplier := &nmodel.ProductSupplier{
			Name:        "Product Supplier 1",
			Address:     "Product Supplier address 1",
			City:        "Product Supplier city 1",
			State:       "Product Supplier state 1",
			Zip:         68215,
			ContactName: "Product Supplier contact name 1",
			Phone:       "081331555666",
			Email:       "productSupplier1@gmail.com",
			URL:         "www.productsupplier1.io",
		}

		mock.ExpectBegin()
		mock.ExpectExec("INSERT").WillReturnError(errors.New("Insert product supplier failed"))
		mock.ExpectCommit()

		_, err := repository.CreateProductSupplier(productSupplier)
		if err == nil {
			t.Errorf("This should not be error")
		}
	})
}
