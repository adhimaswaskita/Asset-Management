package repositories_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nrepo "github.com/adhimaswaskita/AssetManagement/repositories"
	"github.com/jinzhu/gorm"
)

func TestProductType(t *testing.T) {
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

	t.Run("Update Product Type OK", func(t *testing.T) {
		var ID uint = 1
		productType := &nmodel.ProductType{
			Name:        "New product type 1",
			Type:        "asset",
			Category:    "Non-IT",
			Description: "desciption",
		}

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		productTypes, err := repository.UpdateProductType(ID, productType)
		if err != nil {
			t.Errorf("This should not be error but have %v", err)
		}
		if productTypes.Name != "New product type 1" {
			t.Errorf("Product type name should be equal to 'New product type 1' but have %v", productTypes.Name)
		}
	})

	t.Run("Update Product Type NOK", func(t *testing.T) {
		var ID uint = 1
		productType := &nmodel.ProductType{
			Name:        "New product type 1",
			Type:        "asset",
			Category:    "Non-IT",
			Description: "desciption",
		}

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnError(errors.New("Update product type failed"))
		mock.ExpectCommit()

		_, err := repository.UpdateProductType(ID, productType)
		if err == nil {
			t.Errorf("This should be an error")
		}
	})
}
