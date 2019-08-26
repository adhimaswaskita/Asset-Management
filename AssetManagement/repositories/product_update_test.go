package repositories_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nrepo "github.com/adhimaswaskita/AssetManagement/repositories"
	"github.com/jinzhu/gorm"
)

func TestUpdateProduct(t *testing.T) {
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

	t.Run("Update Product OK", func(t *testing.T) {
		var ID uint = 1
		product := &nmodel.Product{
			Name: "New Product  1",
		}

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		products, err := repository.UpdateProduct(ID, product)
		if err != nil {
			t.Errorf("This should not be error but have %v", err)
		}
		if products.Name != "New Product  1" {
			t.Errorf("Product name should be equal to 'New Product  1' but have %v", products.Name)
		}
	})
}
