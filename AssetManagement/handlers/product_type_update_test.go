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

func (m *MockedService) UpdateProductType(ID uint, ProductType *nmodel.ProductType) (*nmodel.ProductType, error) {
	if m.ErrMap["ErrorCreateAllProductType"] {
		return nil, fmt.Errorf("%v", m.ErrStatement)
	}
	producttype := &nmodel.ProductType{
		Model: &gorm.Model{
			ID: 1,
		},
		Name:        "New product type 1",
		Type:        "asset",
		Category:    "Non-IT",
		Description: "desciption",
	}
	return producttype, nil
}

func TestUpdateProductType(t *testing.T) {
	service := &MockedService{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall error",
	}
	handler := nhandler.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/producttype", handler.UpdateProductType).Methods("PUT")
	server := httptest.NewServer(r)
	defer server.Close()

	t.Run("Update Product Type OK", func(t *testing.T) {
		productType := &nmodel.ProductType{
			Name:        "New product type 1",
			Type:        "asset",
			Category:    "Non-IT",
			Description: "desciption",
		}
		jsonProductType, _ := json.Marshal(productType)

		ID := 1
		stringID := strconv.Itoa(ID)

		req, _ := http.NewRequest("PUT", server.URL+"/producttype?ID="+stringID, bytes.NewBuffer(jsonProductType))
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

		productTypeResponse := &nmodel.ProductType{}
		encodeResponse(readBody, productTypeResponse)
		if productTypeResponse.Name != "New product type 1" {
			t.Errorf("productTypeResponse.Name want 'New product type 1' but have %v", productTypeResponse.Name)
		}
	})

	t.Run("Create Product Type NOK", func(t *testing.T) {
		service.ErrMap["ErrorCreateAllProductType"] = true
		productType := &nmodel.ProductType{
			Name:        "New product type 1",
			Type:        "asset",
			Category:    "Non-IT",
			Description: "desciption",
		}
		jsonProductType, _ := json.Marshal(productType)
		req, _ := http.NewRequest("PUT", server.URL+"/producttype", bytes.NewBuffer(jsonProductType))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			t.Errorf("This should be 400 but have: %v", resp.StatusCode)
		}
	})
}
