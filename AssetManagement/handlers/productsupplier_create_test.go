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

func (m *MockedService) CreateProductSupplier(pt *nmodel.ProductSupplier) (*nmodel.ProductSupplier, error) {
	if m.ErrMap["ErrorCreateProductSupplier"] {
		return nil, fmt.Errorf("%v", m.ErrStatement)
	}
	productSupplier := &nmodel.ProductSupplier{
		Model: &gorm.Model{
			ID: 1,
		},
		Name:        "Product Supplier 1",
		Address:     "Product Supplier address 1",
		City:        "Product Supplier city 1",
		State:       "Product Supplier state 1",
		Zip:         68215,
		ContactName: "Product Supplier contact name 1",
		Phone:       6281331555666,
		Email:       "productSupplier1@gmail.com",
		URL:         "www.productsupplier1.io",
	}
	return productSupplier, nil
}

func TestCreateProductSupplier(t *testing.T) {
	service := &MockedService{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}

	handler := nhandler.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/productsupplier", handler.CreateProductSupplier).Methods("POST")
	server := httptest.NewServer(r)
	defer server.Close()

	t.Run("Create Product Supplier OK", func(t *testing.T) {
		productSupplier := &nmodel.ProductSupplier{
			Name:        "Product Supplier 1",
			Address:     "Product Supplier address 1",
			City:        "Product Supplier city 1",
			State:       "Product Supplier state 1",
			Zip:         68215,
			ContactName: "Product Supplier contact name 1",
			Phone:       6281331555666,
			Email:       "productSupplier1@gmail.com",
			URL:         "www.productsupplier1.io",
		}
		jsonProductSupplier, _ := json.Marshal(productSupplier)
		req, _ := http.NewRequest("POST", server.URL+"/productsupplier", bytes.NewBuffer(jsonProductSupplier))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", err)
		}

		readBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("%v", err)
		}

		productSuppliers := nmodel.ProductSupplier{}
		encodeResponse(readBody, &productSuppliers)
		if productSuppliers.ID != 1 {
			t.Errorf("Product Suppliers ID want 1 but have %v", productSuppliers.ID)
		}
	})

	t.Run("Create Product Supplier NOK", func(t *testing.T) {
		service.ErrMap["ErrorCreateProductSupplier"] = true
		productSupplier := &nmodel.ProductSupplier{
			Name:        "Product Supplier 1",
			Address:     "Product Supplier address 1",
			City:        "Product Supplier city 1",
			State:       "Product Supplier state 1",
			Zip:         68215,
			ContactName: "Product Supplier contact name 1",
			Phone:       6281331555666,
			Email:       "productSupplier1@gmail.com",
			URL:         "www.productsupplier1.io",
		}
		jsonProductSupplier, _ := json.Marshal(productSupplier)
		req, _ := http.NewRequest("POST", server.URL+"/productsupplier", bytes.NewBuffer(jsonProductSupplier))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			t.Errorf("This should be 400 but have: %v", resp.StatusCode)
		}
	})
}
