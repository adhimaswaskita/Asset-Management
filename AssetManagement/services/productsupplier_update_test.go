package services_test

import (
	"fmt"
	"testing"

	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nservice "github.com/adhimaswaskita/AssetManagement/services"
	"github.com/jinzhu/gorm"
)

func (m *MockRepository) UpdateProductSupplier(ID uint, ProductSupplier *nmodel.ProductSupplier) (*nmodel.ProductSupplier, error) {
	if m.ErrMap["ErrorUpdateProductSupplier"] {
		return nil, fmt.Errorf("%v", m.ErrStatement)
	}
	productSupplier := &nmodel.ProductSupplier{
		Model: &gorm.Model{
			ID: 1,
		},
		Name:        "New Product Supplier 1",
		Address:     "Product Supplier address 1",
		City:        "Product Supplier city 1",
		State:       "Product Supplier state 1",
		Zip:         68215,
		ContactName: "Product Supplier contact name 1",
		Phone:       "081331555666",
		Email:       "productSupplier1@gmail.com",
		URL:         "www.productsupplier1.io",
	}
	return productSupplier, nil
}

func TestUpdateProductSupplier(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall error",
	}
	service := nservice.Service{
		Repository: repository,
	}
	productSupplier := &nmodel.ProductSupplier{
		Name:        "New Product Supplier 1",
		Address:     "Product Supplier address 1",
		City:        "Product Supplier city 1",
		State:       "Product Supplier state 1",
		Zip:         68215,
		ContactName: "Product Supplier contact name 1",
		Phone:       "081331555666",
		Email:       "productSupplier1@gmail.com",
		URL:         "www.productsupplier1.io",
	}
	t.Run("Update Product Supplier OK", func(t *testing.T) {
		productSuppliers, err := service.UpdateProductSupplier(1, productSupplier)
		if err != nil {
			t.Errorf("This should not be error, but have %v", err)
		}

		if productSuppliers.Name != "New Product Supplier 1" {
			t.Errorf("Product Supplier name want 'New Product Supplier 1' but have %v", productSuppliers.Name)
		}
	})
	t.Run("Update Product Supplier NOK", func(t *testing.T) {
		repository.ErrMap["ErrorUpdateProductSupplier"] = true
		_, err := service.UpdateProductSupplier(1, productSupplier)
		if err == nil {
			t.Errorf("This should be error but no error occured")
		}
	})
}
