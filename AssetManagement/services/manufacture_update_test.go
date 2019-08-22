package services_test

import (
	"fmt"
	"testing"

	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nservice "github.com/adhimaswaskita/AssetManagement/services"
	"github.com/jinzhu/gorm"
)

func (m *MockRepository) UpdateManufacture(ID uint, Manufacture *nmodel.Manufacture) (*nmodel.Manufacture, error) {
	if m.ErrMap["ErrorUpdateManufacture"] {
		return nil, fmt.Errorf("%v", m.ErrStatement)
	}
	manufacture := &nmodel.Manufacture{
		Model: &gorm.Model{
			ID: 1,
		},
		Name:    "New Manufacture 1",
		Comment: "Comment manufacture 1",
	}
	return manufacture, nil
}
func TestUpdateManufacture(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall error",
	}
	service := nservice.Service{
		Repository: repository,
	}
	manufacture := &nmodel.Manufacture{
		Name:    "New manufacture 1",
		Comment: "Comment manufacture 1",
	}
	t.Run("Manufacture Type OK", func(t *testing.T) {
		manufactures, err := service.UpdateManufacture(1, manufacture)
		if err != nil {
			t.Errorf("This should not be error, but have %v", err)
		}

		if manufactures.Name != "New Manufacture 1" {
			t.Errorf("Manufacture name want 'New Manufacture 1' but have %v", manufactures.Name)
		}
	})
	t.Run("Manufacture Type NOK", func(t *testing.T) {
		repository.ErrMap["ErrorUpdateManufacture"] = true
		_, err := service.UpdateManufacture(1, manufacture)
		if err == nil {
			t.Errorf("This should be error but no error occured")
		}
	})
}
