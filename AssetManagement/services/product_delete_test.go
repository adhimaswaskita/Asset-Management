package services_test

import (
	"fmt"
	"testing"

	nservice "github.com/adhimaswaskita/AssetManagement/services"
)

func (m *MockRepository) DeleteProduct(ID uint) error {
	if m.ErrMap["ErrorDeleteProduct"] {
		return fmt.Errorf("%v", m.ErrStatement)
	}

	return nil
}

func TestDeleteProduct(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall error",
	}

	service := nservice.Service{
		Repository: repository,
	}

	t.Run("Delete product  OK", func(t *testing.T) {
		ID := 1
		uintID := uint(ID)
		err := service.DeleteProduct(uintID)
		if err != nil {
			t.Errorf("This should not be error, but have %v", err)
		}
	})
	t.Run("Delete product  NOK", func(t *testing.T) {
		repository.ErrMap["ErrorDeleteProduct"] = true
		ID := 1
		uintID := uint(ID)
		err := service.DeleteProduct(uintID)
		if err == nil {
			t.Errorf("This should be error, but have %v", err)
		}
	})
}
