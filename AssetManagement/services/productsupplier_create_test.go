package services_test

import (
	"fmt"
	"testing"

	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nservice "github.com/adhimaswaskita/AssetManagement/services"
	"github.com/jinzhu/gorm"
)

func (m *MockRepository) CreateProductSupplier(pt *nmodel.ProductSupplier) (*nmodel.ProductSupplier, error) {
	if m.ErrMap["ErrorCreateProductSupplier"] {
		return nil, fmt.Errorf("%v", m.ErrStatement)
	}

	productSupplier := &nmodel.ProductSupplier{
		Model: &gorm.Model{
			ID: 1,
		},
		Name:        "Product Supplier 1",
		Address:     "Product Supplier address 1",
		City:        "Product Supplier city 1",
		State:       "Product Supplier state 1",
		Zip:         68215,
		ContactName: "Product Supplier contact name 1",
		Phone:       6281331555666,
		Email:       "productSupplier1@gmail.com",
		URL:         "www.productsupplier1.io",
	}
	return productSupplier, nil
}
func TestCreateProductSupplier(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}

	service := nservice.Service{
		Repository: repository,
	}

	productSupplier := &nmodel.ProductSupplier{
		Name:        "Product Supplier 1",
		Address:     "Product Supplier address 1",
		City:        "Product Supplier city 1",
		State:       "Product Supplier state 1",
		Zip:         68215,
		ContactName: "Product Supplier contact name 1",
		Phone:       6281331555666,
		Email:       "productSupplier1@gmail.com",
		URL:         "www.productsupplier1.io",
	}

	t.Run("Create Product Supplier OK", func(t *testing.T) {
		productSupplierParam, err := service.CreateProductSupplier(productSupplier)
		if err != nil {
			t.Errorf("This should not be error, but have %v", err)
		}

		if productSupplierParam.ID != 1 {
			t.Errorf("Product Suppliers name want Product Supplier 1 but have %v", productSupplierParam.ID)
		}
	})

	t.Run("Create All Product Supplier NOK", func(t *testing.T) {
		repository.ErrMap["ErrorCreateProductSupplier"] = true
		_, err := service.CreateProductSupplier(productSupplier)
		if err == nil {
			t.Errorf("This should be error")
		}
	})
}
