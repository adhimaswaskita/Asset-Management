package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"

	nhandler "github.com/adhimaswaskita/AssetManagement/handlers"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

func (m *MockedService) CreateProductType(pt *nmodel.ProductType) (*nmodel.ProductType, error) {
	if m.ErrMap["ErrorCreateAllProductType"] {
		return nil, fmt.Errorf("%v", m.ErrStatement)
	}
	producttype := &nmodel.ProductType{
		Model: &gorm.Model{
			ID: 1,
		},
		Name:        "product type 1",
		Type:        "asset",
		Category:    "Non-IT",
		Description: "desciption",
	}
	return producttype, nil
}

func TestCreateProductType(t *testing.T) {
	service := &MockedService{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}

	handler := nhandler.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/producttype", handler.CreateProductType).Methods("POST")
	server := httptest.NewServer(r)
	defer server.Close()

	t.Run("Create Product Type OK", func(t *testing.T) {
		productType := &nmodel.ProductType{
			Name:        "product type 1",
			Type:        "asset",
			Category:    "Non-IT",
			Description: "desciption",
		}
		jsonProductType, _ := json.Marshal(productType)
		req, _ := http.NewRequest("POST", server.URL+"/producttype", bytes.NewBuffer(jsonProductType))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", err)
		}

		readBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("%v", err)
		}

		productTypes := nmodel.ProductType{}
		encodeResponse(readBody, &productTypes)
		if productTypes.ID != 1 {
			t.Errorf("Product types ID want 1 but have %v", productTypes.ID)
		}
	})

	t.Run("Create Product Type NOK", func(t *testing.T) {
		service.ErrMap["ErrorCreateAllProductType"] = true
		productType := &nmodel.ProductType{
			Name:        "product type 1",
			Type:        "asset",
			Category:    "Non-IT",
			Description: "desciption",
		}
		jsonProductType, _ := json.Marshal(productType)
		req, _ := http.NewRequest("POST", server.URL+"/producttype", bytes.NewBuffer(jsonProductType))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			t.Errorf("This should be 400 but have: %v", resp.StatusCode)
		}
	})
}
