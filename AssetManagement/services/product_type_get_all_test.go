package services_test

import (
	"encoding/json"
	"fmt"
	"testing"

	nreadfile "github.com/adhimaswaskita/AssetManagement/lib/readfile"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nservice "github.com/adhimaswaskita/AssetManagement/services"
)

func (r *MockRepository) GetAllProductType() ([]nmodel.ProductType, error) {
	if r.ErrMap["ErrorGetAllProductType"] {
		return nil, fmt.Errorf("%v", r.ErrStatement)
	}
	bProductTypes, err := nreadfile.OpenFile("_fixtures/producttype.json")
	if err != nil {
		return nil, err
	}

	productTypes := []nmodel.ProductType{}
	json.Unmarshal(bProductTypes, &productTypes)

	return productTypes, nil
}

func TestGetAllProductType(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}
	service := nservice.Service{
		Repository: repository,
	}

	t.Run("Get All Product Type", func(t *testing.T) {
		productTypes, err := service.GetAllProductType()
		if err != nil {
			t.Errorf("This should not be error but have %v", err)
		}

		if productTypes[0].ID != 1 {
			t.Errorf("Product Type ID should be 1 but have %v", productTypes[0].ID)
		}
	})

	t.Run("Get All Product Type NOK", func(t *testing.T) {
		repository.ErrMap["ErrorGetAllProductType"] = true
		_, err := service.GetAllProductType()
		if err == nil {
			t.Errorf("This should be error")
		}
	})
}
