package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	nhandler "github.com/adhimaswaskita/AssetManagement/handlers"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

func (m *MockedService) UpdateProductSupplier(ID uint, ProductSupplier *nmodel.ProductSupplier) (*nmodel.ProductSupplier, error) {
	if m.ErrMap["ErrorCreateAllProductSupplier"] {
		return nil, fmt.Errorf("%v", m.ErrStatement)
	}
	productSupplier := &nmodel.ProductSupplier{
		Model: &gorm.Model{
			ID: 1,
		},
		Name:        "New Product Supplier 1",
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

func TestUpdateProductSupplier(t *testing.T) {
	service := &MockedService{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall error",
	}
	handler := nhandler.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/productsupplier", handler.UpdateProductSupplier).Methods("PUT")
	server := httptest.NewServer(r)
	defer server.Close()

	t.Run("Update Product Supplier OK", func(t *testing.T) {
		productSupplier := &nmodel.ProductSupplier{
			Name:        "New Product Supplier 1",
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

		ID := 1
		stringID := strconv.Itoa(ID)

		req, _ := http.NewRequest("PUT", server.URL+"/productsupplier?ID="+stringID, bytes.NewBuffer(jsonProductSupplier))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("This should be 200 but have: %v", resp.StatusCode)
		}
		values := req.URL.Query()["ID"]
		if len(values[0]) < 1 {
			t.Errorf("Parameter want 1 but have %v", values[0])
		}

		if values[0] != "1" {
			t.Errorf("Parameter want 1 but have %v", values[0])
		}

		readBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("%v", err)
		}

		productSupplierResponse := &nmodel.ProductSupplier{}
		encodeResponse(readBody, productSupplierResponse)
		if productSupplierResponse.Name != "New Product Supplier 1" {
			t.Errorf("productSupplierResponse.Name want 'New product Supplier 1' but have %v", productSupplierResponse.Name)
		}
	})

	t.Run("Create Product Supplier NOK", func(t *testing.T) {
		service.ErrMap["ErrorCreateAllProductSupplier"] = true
		productSupplier := &nmodel.ProductSupplier{
			Name:        "New Product Supplier 1",
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
		req, _ := http.NewRequest("PUT", server.URL+"/productsupplier", bytes.NewBuffer(jsonProductSupplier))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			t.Errorf("This should be 400 but have: %v", resp.StatusCode)
		}
	})
}
