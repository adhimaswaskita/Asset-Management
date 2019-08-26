package services_test

import (
	"fmt"
	"testing"

	nservice "github.com/adhimaswaskita/AssetManagement/services"
)

func (m *MockRepository) DeleteProductSupplier(ID uint) error {
	if m.ErrMap["ErrorDeleteProductSupplier"] {
		return fmt.Errorf("%v", m.ErrStatement)
	}

	return nil
}

func TestDeleteProductSupplier(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall error",
	}

	service := nservice.Service{
		Repository: repository,
	}

	t.Run("Delete product supplier OK", func(t *testing.T) {
		ID := 1
		uintID := uint(ID)
		err := service.DeleteProductSupplier(uintID)
		if err != nil {
			t.Errorf("This should not be error, but have %v", err)
		}
	})
	t.Run("Delete product supplier NOK", func(t *testing.T) {
		repository.ErrMap["ErrorDeleteProductSupplier"] = true
		ID := 1
		uintID := uint(ID)
		err := service.DeleteProductSupplier(uintID)
		if err == nil {
			t.Errorf("This should be error, but have %v", err)
		}
	})
}
