package repositories_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nrepo "github.com/adhimaswaskita/AssetManagement/repositories"
	"github.com/jinzhu/gorm"
)

func TestUpdateProductSupplier(t *testing.T) {
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

	t.Run("Update ProductSupplier OK", func(t *testing.T) {
		var ID uint = 1
		productSupplier := &nmodel.ProductSupplier{
			Name: "New Product Supplier 1",
		}

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		productSuppliers, err := repository.UpdateProductSupplier(ID, productSupplier)
		if err != nil {
			t.Errorf("This should not be error but have %v", err)
		}
		if productSuppliers.Name != "New Product Supplier 1" {
			t.Errorf("ProductSupplier name should be equal to 'New Product Supplier 1' but have %v", productSuppliers.Name)
		}
	})

	t.Run("Update Product Supplier NOK", func(t *testing.T) {
		var ID uint = 1
		productSupplier := &nmodel.ProductSupplier{
			Name: "New Product Supplier 1",
		}

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnError(errors.New("Update product supplier failed"))
		mock.ExpectCommit()

		_, err := repository.UpdateProductSupplier(ID, productSupplier)
		if err == nil {
			t.Errorf("This should be an error")
		}
	})
}
