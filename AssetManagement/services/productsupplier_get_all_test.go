package services_test

import (
	"encoding/json"
	"fmt"
	"testing"

	nreadfile "github.com/adhimaswaskita/AssetManagement/lib/readfile"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nservice "github.com/adhimaswaskita/AssetManagement/services"
)

func (r *MockRepository) GetAllProductSupplier() ([]nmodel.ProductSupplier, error) {
	if r.ErrMap["ErrorGetAllProductSupplier"] {
		return nil, fmt.Errorf("%v", r.ErrStatement)
	}
	byteProductSupplier, err := nreadfile.OpenFile("_fixtures/productsupplier.json")
	if err != nil {
		return nil, err
	}

	productSuppliers := []nmodel.ProductSupplier{}
	json.Unmarshal(byteProductSupplier, &productSuppliers)

	return productSuppliers, nil
}

func TestGetAllProductSupplier(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}
	service := nservice.Service{
		Repository: repository,
	}

	t.Run("Get All Product Supplier", func(t *testing.T) {
		productSuppliers, err := service.GetAllProductSupplier()
		if err != nil {
			t.Errorf("This should not be error but have %v", err)
		}

		if productSuppliers[0].ID != 1 {
			t.Errorf("Product Supplier ID should be 1 but have %v", productSuppliers[0].ID)
		}
	})

	t.Run("Get All Product Supplier NOK", func(t *testing.T) {
		repository.ErrMap["ErrorGetAllProductSupplier"] = true
		_, err := service.GetAllProductSupplier()
		if err == nil {
			t.Errorf("This should be error")
		}
	})
}
