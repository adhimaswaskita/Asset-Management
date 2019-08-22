package services_test

import (
	"encoding/json"
	"fmt"
	"testing"

	nreadfile "github.com/adhimaswaskita/AssetManagement/lib/readfile"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nservice "github.com/adhimaswaskita/AssetManagement/services"
)

func (r *MockRepository) GetAllManufacture() ([]nmodel.Manufacture, error) {
	if r.ErrMap["ErrorGetAllManufacture"] {
		return nil, fmt.Errorf("%v", r.ErrStatement)
	}
	bManufacture, err := nreadfile.OpenFile("_fixtures/manufacture.json")
	if err != nil {
		return nil, err
	}

	manufactures := []nmodel.Manufacture{}
	json.Unmarshal(bManufacture, &manufactures)

	return manufactures, nil
}

func TestGetAllManufacture(t *testing.T) {
	repository := &MockRepository{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}
	service := nservice.Service{
		Repository: repository,
	}

	t.Run("Get All Manufacture", func(t *testing.T) {
		manufactures, err := service.GetAllManufacture()
		if err != nil {
			t.Errorf("This should not be error but have %v", err)
		}

		if manufactures[0].ID != 1 {
			t.Errorf("Manufacture ID should be 1 but have %v", manufactures[0].ID)
		}
	})

	t.Run("Get All Manufacture NOK", func(t *testing.T) {
		repository.ErrMap["ErrorGetAllManufacture"] = true
		_, err := service.GetAllManufacture()
		if err == nil {
			t.Errorf("This should be error")
		}
	})
}
