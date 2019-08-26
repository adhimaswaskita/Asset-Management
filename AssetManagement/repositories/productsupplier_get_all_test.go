package repositories_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	nrepo "github.com/adhimaswaskita/AssetManagement/repositories"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestGetAllProductSupplier(t *testing.T) {
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

	t.Run("Testing Get All Product Supplier OK", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Name", "Address", "City", "State", "Country", "Zip", "ContactName", "Phone", "Email", "URL"}).AddRow("1", "2019-07-31T13:38:25+07:00", "2019-07-31T13:38:25+07:00", nil, "Product Supplier 1", "Perum peruri", "Bondowoso", "East Java", "Indonesia", 68215, "product supplier", 81331555666, "productsupplier@gmail.com", "www.productsupplier.com")

		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		productSuppliers, err := repository.GetAllProductSupplier()
		if err != nil {
			t.Errorf("This should not be error")
		}
		if len(productSuppliers) <= 0 {
			t.Errorf("This should not be 0")
		}
	})
}
