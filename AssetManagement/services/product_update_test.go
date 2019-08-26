package services_test

import (
	"fmt"
	"testing"

	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nservice "github.com/adhimaswaskita/AssetManagement/services"
	"github.com/jinzhu/gorm"
)

func (m *MockRepository) UpdateProduct(ID uint, Product *nmodel.Product) (*nmodel.Product, error) {
	if m.ErrMap["ErrorUpdateProduct"] {
		return nil, fmt.Errorf("%v", m.ErrStatement)
	}
	product := &nmodel.Product{
		Model: &gorm.Model{
			ID: 1,
		},
		Name:              "New Product 1",
		ManufactureID:     1,
		ProductTypeID:     1,
		ProductSupplierID: 1,
		Manufacturer:      "Product 1 manufacturer",
		PartNo:            19,
	}
	return product, nil
}

func TestUpdateProduct(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall error",
	}
	service := nservice.Service{
		Repository: repository,
	}
	product := &nmodel.Product{
		Name:              "New Product 1",
		ManufactureID:     1,
		ProductTypeID:     1,
		ProductSupplierID: 1,
		Manufacturer:      "Product 1 manufacturer",
		PartNo:            19,
	}
	t.Run("Update Product OK", func(t *testing.T) {
		products, err := service.UpdateProduct(1, product)
		if err != nil {
			t.Errorf("This should not be error, but have %v", err)
		}

		if products.Name != "New Product 1" {
			t.Errorf("Product  name want 'New Product 1' but have %v", products.Name)
		}
	})
	t.Run("Update Product  NOK", func(t *testing.T) {
		repository.ErrMap["ErrorUpdateProduct"] = true
		_, err := service.UpdateProduct(1, product)
		if err == nil {
			t.Errorf("This should be error but no error occured")
		}
	})
}
