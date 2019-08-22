package services_test

import (
	"fmt"
	"testing"

	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nservice "github.com/adhimaswaskita/AssetManagement/services"
	"github.com/jinzhu/gorm"
)

func (m *MockRepository) CreateManufacture(pt *nmodel.Manufacture) (*nmodel.Manufacture, error) {
	if m.ErrMap["ErrorCreateManufacture"] {
		return nil, fmt.Errorf("%v", m.ErrStatement)
	}

	manufacture := &nmodel.Manufacture{
		Model: &gorm.Model{
			ID: 1,
		},
		Name:    "manufacture 1",
		Comment: "Comment",
	}
	return manufacture, nil
}
func TestCreateManufacture(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}

	service := nservice.Service{
		Repository: repository,
	}

	manufacture := &nmodel.Manufacture{
		Name:    "manufacture 1",
		Comment: "Comment",
	}

	t.Run("Create Product Type OK", func(t *testing.T) {
		manufactures, err := service.CreateManufacture(manufacture)
		if err != nil {
			t.Errorf("This should not be error, but have %v", err)
		}

		if manufactures.ID != 1 {
			t.Errorf("Product Types name want product type 1 but have %v", manufactures.ID)
		}
	})

	t.Run("Create All Product Type NOK", func(t *testing.T) {
		repository.ErrMap["ErrorCreateManufacture"] = true
		_, err := service.CreateManufacture(manufacture)
		if err == nil {
			t.Errorf("This should be error")
		}
	})
}
