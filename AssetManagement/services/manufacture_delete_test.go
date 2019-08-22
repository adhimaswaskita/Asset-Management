package services_test

import (
	"fmt"
	"testing"

	nservice "github.com/adhimaswaskita/AssetManagement/services"
)

func (m *MockRepository) DeleteManufacture(ID uint) error {
	if m.ErrMap["ErrorDeleteProductType"] {
		return fmt.Errorf("%v", m.ErrStatement)
	}

	return nil
}

func TestDeleteManufacture(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall error",
	}

	service := nservice.Service{
		Repository: repository,
	}

	ID := uint(1)
	t.Run("Delete manufacture OK", func(t *testing.T) {
		err := service.DeleteManufacture(ID)
		if err != nil {
			t.Errorf("This should not be error, but have %v", err)
		}
	})
	t.Run("Delete manufacture NOK", func(t *testing.T) {
		repository.ErrMap["ErrorDeleteProductType"] = true
		err := service.DeleteManufacture(ID)
		if err == nil {
			t.Errorf("This should be error, but have %v", err)
		}
	})
}
