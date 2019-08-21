package services_test

import (
	"fmt"
	"testing"

	nservice "github.com/adhimaswaskita/AssetManagement/services"
)

func (m *MockRepository) DeleteProductType(ID int) error {
	if m.ErrMap["ErrorCreateAllProductType"] {
		return fmt.Errorf("%v", m.ErrStatement)
	}

	return nil
}

func TestDeleteProductType(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall error",
	}

	service := nservice.Service{
		Repository: repository,
	}

	t.Run("Delete product type OK", func(t *testing.T) {
		ID := 1
		err := service.DeleteProductType(ID)
		if err != nil {
			t.Errorf("This should not be error, but have %v", err)
		}
	})
	t.Run("Delete product type NOK", func(t *testing.T) {
		repository.ErrMap["ErrorCreateAllProductType"] = true
		ID := 1
		err := service.DeleteProductType(ID)
		if err == nil {
			t.Errorf("This should be error, but have %v", err)
		}
	})
}
