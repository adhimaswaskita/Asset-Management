package repositories_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	nrepo "github.com/adhimaswaskita/netmonk-asset-management/repositories"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestGetAllProductType(t *testing.T) {
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

	t.Run("Testing Find Get All Product Type OK", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Name", "Type", "Category", "Description"}).AddRow("1", "2019-07-31T13:38:25+07:00", "2019-07-31T13:38:25+07:00", nil, "Product Type 1", "Asset", "IT", "Product type 1 description")

		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		productTypes, err := repository.GetAllProductType()
		if err != nil {
			t.Errorf("This should not be error")
		}
		if len(productTypes) <= 0 {
			t.Errorf("This should not be 0")
		}
	})
}
