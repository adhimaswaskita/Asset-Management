package services_test

import (
	"fmt"
	"testing"

	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nservice "github.com/adhimaswaskita/AssetManagement/services"
	"github.com/jinzhu/gorm"
)

func (m *MockRepository) UpdateProductType(ID uint, ProductType *nmodel.ProductType) (*nmodel.ProductType, error) {
	if m.ErrMap["ErrorUpdateProductType"] {
		return nil, fmt.Errorf("%v", m.ErrStatement)
	}
	producttype := &nmodel.ProductType{
		Model: &gorm.Model{
			ID: 1,
		},
		Name:        "New product type 1",
		Type:        "asset",
		Category:    "Non-IT",
		Description: "desciption",
	}
	return producttype, nil
}

func TestUpdateProductType(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall error",
	}
	service := nservice.Service{
		Repository: repository,
	}
	productType := &nmodel.ProductType{
		Name:        "New product type 1",
		Type:        "asset",
		Category:    "Non-IT",
		Description: "desciption",
	}
	t.Run("Update Product Type OK", func(t *testing.T) {
		productTypes, err := service.UpdateProductType(1, productType)
		if err != nil {
			t.Errorf("This should not be error, but have %v", err)
		}

		if productTypes.Name != "New product type 1" {
			t.Errorf("Product Type name want 'New product type 1' but have %v", productTypes.Name)
		}
	})
	t.Run("Update Product Type NOK", func(t *testing.T) {
		repository.ErrMap["ErrorUpdateProductType"] = true
		_, err := service.UpdateProductType(1, productType)
		if err == nil {
			t.Errorf("This should be error but no error occured")
		}
	})
}
