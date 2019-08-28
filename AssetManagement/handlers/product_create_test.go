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

func (m *MockedService) CreateProduct(pt *nmodel.Product) (*nmodel.Product, error) {
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
		PartNo:            19,
	}
	return product, nil
}

func TestCreateProduct(t *testing.T) {
	service := &MockedService{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}

	handler := nhandler.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/product", handler.CreateProduct).Methods("POST")
	server := httptest.NewServer(r)
	defer server.Close()

	t.Run("Create Product  OK", func(t *testing.T) {
		product := &nmodel.Product{
			Name:              "Product 1",
			ManufactureID:     1,
			ProductTypeID:     1,
			ProductSupplierID: 1,
			PartNo:            19,
		}
		jsonProduct, _ := json.Marshal(product)
		req, _ := http.NewRequest("POST", server.URL+"/product", bytes.NewBuffer(jsonProduct))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", err)
		}

		readBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("%v", err)
		}

		products := nmodel.Product{}
		encodeResponse(readBody, &products)
		if products.ID != 1 {
			t.Errorf("Product s ID want 1 but have %v", products.ID)
		}
	})

	t.Run("Create Product  NOK", func(t *testing.T) {
		service.ErrMap["ErrorCreateProduct"] = true
		product := &nmodel.Product{
			Name:              "Product 1",
			ManufactureID:     1,
			ProductTypeID:     1,
			ProductSupplierID: 1,
			PartNo:            19,
		}
		jsonProduct, _ := json.Marshal(product)
		req, _ := http.NewRequest("POST", server.URL+"/product", bytes.NewBuffer(jsonProduct))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			t.Errorf("This should be 400 but have: %v", resp.StatusCode)
		}
	})
}
