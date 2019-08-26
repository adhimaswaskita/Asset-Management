package services_test

import (
	"encoding/json"
	"fmt"
	"testing"

	nreadfile "github.com/adhimaswaskita/AssetManagement/lib/readfile"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nservice "github.com/adhimaswaskita/AssetManagement/services"
)

func (r *MockRepository) GetAllProduct() ([]nmodel.Product, error) {
	if r.ErrMap["ErrorGetAllProduct"] {
		return nil, fmt.Errorf("%v", r.ErrStatement)
	}
	byteProduct, err := nreadfile.OpenFile("_fixtures/product.json")
	if err != nil {
		return nil, err
	}

	products := []nmodel.Product{}
	json.Unmarshal(byteProduct, &products)

	return products, nil
}

func TestGetAllProduct(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}
	service := nservice.Service{
		Repository: repository,
	}

	t.Run("Get All Product ", func(t *testing.T) {
		products, err := service.GetAllProduct()
		if err != nil {
			t.Errorf("This should not be error but have %v", err)
		}

		if products[0].ID != 1 {
			t.Errorf("Product  ID should be 1 but have %v", products[0].ID)
		}
	})

	t.Run("Get All Product  NOK", func(t *testing.T) {
		repository.ErrMap["ErrorGetAllProduct"] = true
		_, err := service.GetAllProduct()
		if err == nil {
			t.Errorf("This should be error")
		}
	})
}
