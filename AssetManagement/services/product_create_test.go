package services_test

import (
	"fmt"
	"testing"

	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nservice "github.com/adhimaswaskita/AssetManagement/services"
	"github.com/jinzhu/gorm"
)

func (m *MockRepository) CreateProduct(pt *nmodel.Product) (*nmodel.Product, error) {
	if m.ErrMap["ErrorCreateProduct"] {
		return nil, fmt.Errorf("%v", m.ErrStatement)
	}

	product := &nmodel.Product{
		Model: &gorm.Model{
			ID: 1,
		},
		Name:              "Product 1",
		ManufactureID:     1,
		ProductTypeID:     1,
		ProductSupplierID: 1,
		Manufacturer:      "Product 1 manufacturer",
		PartNo:            19,
	}
	return product, nil
}
func TestCreateProduct(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}

	service := nservice.Service{
		Repository: repository,
	}

	product := &nmodel.Product{
		Name:              "Product 1",
		ManufactureID:     1,
		ProductTypeID:     1,
		ProductSupplierID: 1,
		Manufacturer:      "Product 1 manufacturer",
		PartNo:            19,
	}

	t.Run("Create Product  OK", func(t *testing.T) {
		productParam, err := service.CreateProduct(product)
		if err != nil {
			t.Errorf("This should not be error, but have %v", err)
		}

		if productParam.ID != 1 {
			t.Errorf("Product s name want Product  1 but have %v", productParam.ID)
		}
	})

	t.Run("Create All Product  NOK", func(t *testing.T) {
		repository.ErrMap["ErrorCreateProduct"] = true
		_, err := service.CreateProduct(product)
		if err == nil {
			t.Errorf("This should be error")
		}
	})
}
