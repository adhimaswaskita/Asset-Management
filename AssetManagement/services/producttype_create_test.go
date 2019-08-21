package services_test

import (
	"fmt"
	"testing"

	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nservice "github.com/adhimaswaskita/AssetManagement/services"
	"github.com/jinzhu/gorm"
)

func (m *MockRepository) CreateProductType(pt *nmodel.ProductType) (*nmodel.ProductType, error) {
	if m.ErrMap["ErrorCreateAllProductType"] {
		return nil, fmt.Errorf("%v", m.ErrStatement)
	}

	producttype := &nmodel.ProductType{
		Model: &gorm.Model{
			ID: 1,
		},
		Name:        "product type 1",
		Type:        "asset",
		Category:    "Non-IT",
		Description: "desciption",
	}
	return producttype, nil
}
func TestCreateProductType(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}

	service := nservice.Service{
		Repository: repository,
	}

	productType := &nmodel.ProductType{
		Name:        "product type 1",
		Type:        "asset",
		Category:    "Non-IT",
		Description: "desciption",
	}

	t.Run("Create Product Type OK", func(t *testing.T) {
		productTypes, err := service.CreateProductType(productType)
		if err != nil {
			t.Errorf("This should not be error, but have %v", err)
		}

		if productTypes.ID != 1 {
			t.Errorf("Product Types name want product type 1 but have %v", productTypes.ID)
		}
	})

	t.Run("Create All Product Type NOK", func(t *testing.T) {
		repository.ErrMap["ErrorCreateAllProductType"] = true
		_, err := service.CreateProductType(productType)
		if err == nil {
			t.Errorf("This should be error")
		}
	})
}
