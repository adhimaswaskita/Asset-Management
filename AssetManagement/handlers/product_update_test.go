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

func (m *MockedService) UpdateProduct(ID uint, Product *nmodel.Product) (*nmodel.Product, error) {
	if m.ErrMap["ErrorCreateAllProduct"] {
		return nil, fmt.Errorf("%v", m.ErrStatement)
	}
	product := &nmodel.Product{
		Model: &gorm.Model{
			ID: 1,
		},
		Name:              "New Product 1",
		ManufactureID:     1,
		ProductTypeID:     1,
		ProductSupplierID: 1,
		Manufacturer:      "New Product 1 manufacturer",
		PartNo:            19,
	}
	return product, nil
}

func TestUpdateProduct(t *testing.T) {
	service := &MockedService{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall error",
	}
	handler := nhandler.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/product", handler.UpdateProduct).Methods("PUT")
	server := httptest.NewServer(r)
	defer server.Close()

	t.Run("Update Product  OK", func(t *testing.T) {
		product := &nmodel.Product{
			Name:              "New Product 1",
			ManufactureID:     1,
			ProductTypeID:     1,
			ProductSupplierID: 1,
			Manufacturer:      "New Product 1 manufacturer",
			PartNo:            19,
		}
		jsonProduct, _ := json.Marshal(product)

		ID := 1
		stringID := strconv.Itoa(ID)

		req, _ := http.NewRequest("PUT", server.URL+"/product?ID="+stringID, bytes.NewBuffer(jsonProduct))
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

		productResponse := &nmodel.Product{}
		encodeResponse(readBody, productResponse)
		if productResponse.Name != "New Product 1" {
			t.Errorf("productResponse.Name want 'New product  1' but have %v", productResponse.Name)
		}
	})

	t.Run("Create Product  NOK", func(t *testing.T) {
		service.ErrMap["ErrorCreateAllProduct"] = true
		product := &nmodel.Product{
			Name:              "New Product 1",
			ManufactureID:     1,
			ProductTypeID:     1,
			ProductSupplierID: 1,
			Manufacturer:      "New Product 1 manufacturer",
			PartNo:            19,
		}
		jsonProduct, _ := json.Marshal(product)
		req, _ := http.NewRequest("PUT", server.URL+"/product", bytes.NewBuffer(jsonProduct))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			t.Errorf("This should be 400 but have: %v", resp.StatusCode)
		}
	})
}
